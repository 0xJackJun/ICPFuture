use ic_cdk_macros::*;
use ic_cdk::{api, storage};
use candid::{candid_method, Principal, Deserialize, CandidType};
use ic_cron::types::{Iterations,SchedulingInterval};

struct Counter(Principal);
struct Owner(Principal);

#[derive(CandidType, Deserialize)]
struct UserData {
    canister_id: String,
    email: String,
    threshold: String,
  }

impl Default for Counter {
    fn default() -> Self {
        Counter(Principal::anonymous())
    }
}

impl Default for Owner {
    fn default() -> Self {
        Owner(Principal::anonymous())
    }
}

#[init]
#[candid_method(init)]
fn init() {
    let caller = ic_cdk::caller();
    let owner = storage::get_mut::<Owner>();
    *owner = Owner(caller);
}

#[update(name = "setCounter")]
#[candid_method(update, rename = "setCounter")]
fn set_counter(counter: Principal) {
    let caller = ic_cdk::caller();
    let owner = storage::get::<Owner>();
    api::print(owner.0.to_text());
    assert_eq!(caller, owner.0);

    let _counter = storage::get_mut::<Counter>();
    _counter.0 = counter;
}


// cross-canister calling
#[update]
#[candid_method]
async fn get(duration_nano: u64) -> Vec<UserData> {
    let counter = storage::get::<Counter>();
    let result: Result<(Vec<UserData>,), _> = api::call::call(counter.0, "getUser", ()).await;
    cron_enqueue(
        TaskKind::ReadData as u8,
        String::from("sweetie"), 
        SchedulingInterval {
            duration_nano,
            iterations: Iterations::Infinite,
        },
    );
    result.unwrap().0
}


#[cfg(any(target_arch = "wasm32", test))]
fn main() {}

#[cfg(not(any(target_arch = "wasm32", test)))]
fn main() {
    candid::export_service!();
    std::print!("{}", __export_service());
}

ic_cron::implement_cron!();

ic_cron::u8_enum! {
    pub enum TaskKind {
        ReadData,
    }
}

// the function do cron job, which can periodical tasks.
#[heartbeat]
fn heartbeat() {
    for task in cron_ready_tasks() {
        match task.get_kind().try_into() {
            Ok(TaskKind::ReadData) => {
                get(360_000_000_000); //call get function per hour
            },
            Err(_) => ic_cdk::print("Error: Invalid cron task handler"),
        };   
    }
}
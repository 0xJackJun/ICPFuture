import Principal "mo:base/Principal";

module {
  public type CanisterId = Principal;
  
  public type UserId = Principal;

  public type DefiniteCanisterSettings = {
    freezing_threshold : Nat;
    controllers : [Principal];
    memory_allocation : Nat;
    compute_allocation : Nat;
  };

  public type CanisterStatus = {
     status : { #stopped; #stopping; #running };
     memory_size : Nat;
     cycles : Nat;
     settings : DefiniteCanisterSettings;
     module_hash : ?[Nat8];
  };
  
  public type UserData = {
    canister_id: Text;
    email: Text;
    threshold: Text;
  };
};

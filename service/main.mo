import Blackhole "canister:blackhole";
import Types "./types";
import Database "./database";

actor Cyclebalance{
    var data: Database.Data = Database.Data();
    type UserData = Types.UserData;
    type canister_id = Principal;
    type Status = Types.CanisterStatus;
    
    public func getCycle(request : { canister_id : canister_id }) : async Nat {
        var status: Status = await Blackhole.canister_status(request);
        var cycles: Nat = status.cycles;
        return cycles;
    };

    public shared(msg) func create(userData: UserData): async() {
        data.createOne(userData);
    };

    public func getUser(): async[UserData]{
        data.get()
    };

    // public query func get(userIds: [UserId]): [UserData]{
    //     data.getMany(userIds)
    // };
};
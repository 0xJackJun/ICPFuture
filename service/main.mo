import Blackhole "canister:blackhole";
import Types "./types";

actor Cyclebalance{
    type canister_id = Principal;
    
    type Status = Types.CanisterStatus;
    
    public func getStatus(request : { canister_id : canister_id }) : async Status {
        await Blackhole.canister_status(request)
    }
}
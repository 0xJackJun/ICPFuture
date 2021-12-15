import Types "./types";
import Array "mo:base/Array";
import Iter "mo:base/Iter";
import Principal "mo:base/Principal";

module {
  type UserData = Types.UserData;
  type UserId = Types.UserId;

  public class Data(){
    // the "database" is just a local array
    var array: [UserData] = [];

    public func createOne(userData: UserData) {
        array := Array.append<UserData>(array, [userData]);
    };

    public func get(): [UserData]{
        array
    };

    // public func getMany(userIds: [UserId]): [UserData]{
    //     func getUserData(userId: UserId): UserData{
    //         Option.unwrap<UserData>(hashMap.get(userId))
    //     };
    //     Array.map<UserId, UserData>(userId,getUserData)
    // };
    
    // public func get():[UserData]{
    //     return hashMap.v
    // } 

    // public func makeUserData(userId: UserId, userData: UserData): UserData{
    //     {
    //         id = userId;
    //         canister_id = userData.canister_id;
    //         email = userData.email;
    //         threshold = userData.threshold;
    //     }
    // };
  };
};
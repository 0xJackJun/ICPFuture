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
    var count: Nat = 0;
    public func createOne(userData: UserData) {
        count := count + 1;
        array := Array.append<UserData>(array, [userData]);
    };

    public func get(): [UserData]{
        array
    };

    public func getCount(): Nat{
        count
    }
  };
};
import type { Principal } from '@dfinity/agent';
export interface UserData {
  'threshold' : string,
  'canister_id' : string,
  'email' : string,
};
export default interface _SERVICE {
  'get' : () => Promise<Array<UserData>>,
  'setCounter' : (arg_0: Principal) => Promise<undefined>,
};
import type { Principal } from '@dfinity/agent';
export type UserData = UserData_2;
export interface UserData_2 {
  'threshold' : string,
  'canister_id' : string,
  'email' : string,
};
export type canister_id = Principal;
export default interface _SERVICE {
  'create' : (arg_0: UserData) => Promise<undefined>,
  'getCycle' : (arg_0: { 'canister_id' : canister_id }) => Promise<bigint>,
  'getUser' : () => Promise<Array<UserData>>,
};
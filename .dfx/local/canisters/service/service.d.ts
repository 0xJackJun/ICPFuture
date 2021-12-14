import type { Principal } from '@dfinity/agent';
export interface CanisterStatus {
  'status' : { 'stopped' : null } |
    { 'stopping' : null } |
    { 'running' : null },
  'memory_size' : bigint,
  'cycles' : bigint,
  'settings' : DefiniteCanisterSettings,
  'module_hash' : [] | [Array<number>],
};
export interface DefiniteCanisterSettings {
  'freezing_threshold' : bigint,
  'controllers' : Array<Principal>,
  'memory_allocation' : bigint,
  'compute_allocation' : bigint,
};
export type Status = CanisterStatus;
export type canister_id = Principal;
export default interface _SERVICE {
  'getStatus' : (arg_0: { 'canister_id' : canister_id }) => Promise<Status>,
};
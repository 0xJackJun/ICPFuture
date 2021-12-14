export default ({ IDL }) => {
  const canister_id = IDL.Principal;
  const DefiniteCanisterSettings = IDL.Record({
    'freezing_threshold' : IDL.Nat,
    'controllers' : IDL.Vec(IDL.Principal),
    'memory_allocation' : IDL.Nat,
    'compute_allocation' : IDL.Nat,
  });
  const CanisterStatus = IDL.Record({
    'status' : IDL.Variant({
      'stopped' : IDL.Null,
      'stopping' : IDL.Null,
      'running' : IDL.Null,
    }),
    'memory_size' : IDL.Nat,
    'cycles' : IDL.Nat,
    'settings' : DefiniteCanisterSettings,
    'module_hash' : IDL.Opt(IDL.Vec(IDL.Nat8)),
  });
  const Status = CanisterStatus;
  return IDL.Service({
    'getStatus' : IDL.Func(
        [IDL.Record({ 'canister_id' : canister_id })],
        [Status],
        [],
      ),
  });
};
export const init = ({ IDL }) => { return []; };
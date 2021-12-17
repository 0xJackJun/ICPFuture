export default ({ IDL }) => {
  const UserData_2 = IDL.Record({
    'threshold' : IDL.Text,
    'canister_id' : IDL.Text,
    'email' : IDL.Text,
  });
  const UserData = UserData_2;
  const canister_id = IDL.Principal;
  return IDL.Service({
    'create' : IDL.Func([UserData], [], []),
    'getCount' : IDL.Func([], [IDL.Nat], []),
    'getCycle' : IDL.Func(
        [IDL.Record({ 'canister_id' : canister_id })],
        [IDL.Nat],
        [],
      ),
    'getUser' : IDL.Func([], [IDL.Vec(UserData)], []),
  });
};
export const init = ({ IDL }) => { return []; };
export default ({ IDL }) => {
  const UserData = IDL.Record({
    'threshold' : IDL.Text,
    'canister_id' : IDL.Text,
    'email' : IDL.Text,
  });
  return IDL.Service({
    'get' : IDL.Func([], [IDL.Vec(UserData)], []),
    'setCounter' : IDL.Func([IDL.Principal], [], []),
  });
};
export const init = ({ IDL }) => { return []; };
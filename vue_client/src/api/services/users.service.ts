import { useQuery } from '@vue/apollo-composable';
import { LIST_USERS_QUERY } from '../graphql/users.qm';

export const listUsers = () => {
  return useQuery(LIST_USERS_QUERY);
};

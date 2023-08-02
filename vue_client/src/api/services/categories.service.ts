import {
  provideApolloClient,
  useMutation,
  useQuery,
} from '@vue/apollo-composable';
import { apolloClient } from '../apolloClient';
import {
  CREATE_CATEGORY_MUTATION,
  LABLE_AVAILABILITY_QUERY,
  LIST_CATEGORIES_QUERY,
} from '../graphql/catrgories.qm';

export const listCategories = () => {
  return useQuery(LIST_CATEGORIES_QUERY);
};

export const checkLableAvailability = (lable: string) => {
  return new Promise<boolean>((resolve, reject) => {
    provideApolloClient(apolloClient)(() => {
      const { onError, onResult } = useQuery(LABLE_AVAILABILITY_QUERY, {
        lable: lable,
      });
      onResult(({ data }) => {
        resolve(data.checkLableAvailability);
      });
      onError(() => {
        reject();
      });
    });
  });
};

export const createCategoryMutation = (
  createCategoryRequest: ICreateCategory
) => {
  return useMutation(CREATE_CATEGORY_MUTATION, {
    variables: { input: createCategoryRequest },
  });
};

export interface ICreateCategory {
  lable: string;
  description: string;
}

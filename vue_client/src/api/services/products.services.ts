import { useMutation, useQuery } from '@vue/apollo-composable';
import {
  CREATE_PRODUCT_MUTATION,
  LIST_PRODUCTS_QUERY,
} from '../graphql/products.qm';

export const listProductsQuery = () => {
  return useQuery(LIST_PRODUCTS_QUERY);
};

export const createProductMutation = (createProductRequest: ICreateProduct) => {
  return useMutation(CREATE_PRODUCT_MUTATION, {
    variables: { input: createProductRequest },
  });
};

type Nullable<T> = T | null;

export interface ICreateProduct {
  name_product: string;
  description_product: string;
  price_product: number;
  quantity_product: number;
  is_available: boolean;
  product_image: Nullable<string>;
  category_id: Nullable<number>;
}

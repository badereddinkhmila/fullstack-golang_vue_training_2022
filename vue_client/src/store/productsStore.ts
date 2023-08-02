import { apolloClient } from '@/api/apolloClient';
import {
  createCategoryMutation,
  ICreateCategory,
  listCategories,
} from '@/api/services/categories.service';
import {
  createProductMutation,
  ICreateProduct,
  listProductsQuery,
} from '@/api/services/products.services';
import { ApolloError } from '@apollo/client';
import { provideApolloClient } from '@vue/apollo-composable';
import { defineStore } from 'pinia';

const useProductsStore = defineStore({
  id: 'products',
  state: (): ProductsState => ({
    products: [],
    categories: [],
    isLoading: false,
    errors: null,
  }),
  actions: {
    getAllProducts() {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const { onResult, onError } = listProductsQuery();
          onError((err) => {
            this.isLoading = false;
            this.errors = err;
            reject(err);
          });
          onResult(({ data }) => {
            if (data.products != null) {
              this.products = data.products;
              this.isLoading = false;
              resolve();
            }
          });
        });
      });
    },
    getCategories() {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const { onResult, onError } = listCategories();
          onError((err) => {
            this.categories = [];
            this.isLoading = false;
            this.errors = err;
            reject(err);
          });
          onResult(({ data }) => {
            if (data.categories != null) {
              this.categories = data.categories;
              this.isLoading = false;
              resolve();
            }
          });
        });
      });
    },
    createProduct(product: ICreateProduct) {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const {
            mutate: createProduct,
            onDone,
            onError,
          } = createProductMutation(product);
          createProduct();
          onError((err) => {
            this.isLoading = false;
            reject(err);
          });
          onDone(({ data }) => {
            this.isLoading = false;
            console.log(data);
            resolve();
          });
        });
      });
    },
    createCategory(category: ICreateCategory) {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const {
            mutate: createCategory,
            onDone,
            onError,
          } = createCategoryMutation(category);
          createCategory();
          onError((err) => {
            this.isLoading = false;
            reject(err);
          });
          onDone(({ data }) => {
            this.isLoading = false;
            console.log(data);
            resolve();
          });
        });
      });
    },
    filterProductById(id: string) {
      return this.products.filter((product) => `${product.id}` === id);
    },
  },
  getters: {},
  persist: {
    enabled: true,
    strategies: [
      {
        key: 'products',
        storage: localStorage,
      },
    ],
  },
});
export default useProductsStore;

interface ProductsState {
  products: Product[];
  categories: Category[];
  isLoading: boolean;
  errors: ApolloError | null;
}

type Nullable<T> = T | null;

export interface Product {
  id: number;
  nameProduct: string;
  descriptionProduct: string;
  priceProduct: number;
  quantity: number;
  isAvailable: boolean;
  imageProduct: Nullable<string>;
  categoryId: Nullable<number>;
}

export interface Category {
  id: number;
  lable: string;
  description: string;
}

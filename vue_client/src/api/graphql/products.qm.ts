import gql from 'graphql-tag';

export const CREATE_PRODUCT_MUTATION = gql`
  mutation createProduct($input: CreateProductInput!) {
    createProduct(product: $input) {
      id
      price_product
      name_product
      description_product
      category_id
      product_image
      quantity_product
      is_available
    }
  }
`;

export const LIST_PRODUCTS_QUERY = gql`
  query listProducts {
    products {
      id
      price_product
      name_product
      description_product
      product_image
      quantity_product
      is_available
      created_at
    }
  }
`;

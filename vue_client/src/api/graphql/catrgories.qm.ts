import gql from 'graphql-tag';

export const CREATE_CATEGORY_MUTATION = gql`
  mutation createCategory($input: CreateCategoryInput!) {
    createCategory(category: $input) {
      id
      lable
      description
      created_by {
        id
        username
      }
      created_at
    }
  }
`;

export const UPDATE_CATEGORY_MUTATION = gql`
  mutation updateCategory($input: UpdateCategoryInput!) {
    updateCategory(category: $input) {
      id
      lable
      description
      created_by {
        id
        username
      }
      created_at
    }
  }
`;

//export const DELETE_CATEGORY_MUTATION = gql``;

export const LIST_CATEGORIES_QUERY = gql`
  query categories {
    categories {
      id
      lable
      description
      created_at
    }
  }
`;

export const LABLE_AVAILABILITY_QUERY = gql`
  query lableAvailability($lable: String!) {
    checkLableAvailability(lable: $lable)
  }
`;

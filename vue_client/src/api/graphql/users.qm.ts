import gql from 'graphql-tag';

export const LOGIN_USER_QUERY = gql`
  query loginUser($input: LoginInput!) {
    loginUser(input: $input) {
      access_token
      refresh_token
      current_user {
        id
        username
        email
        role
      }
    }
  }
`;

export const USERNAME_AVAILABILITY_QUERY = gql`
  query usernameAvailability($username: String!) {
    checkUsernameAvailability(username: $username)
  }
`;

export const EMAIL_AVAILABILITY_QUERY = gql`
  query emailAvailability($email: String!) {
    checkEmailAvailability(email: $email)
  }
`;

export const REFRESH_USER_TOKEN = gql`
  query refreshToken($token: String!) {
    refreshToken(token: $token) {
      access_token
      refresh_token
      current_user {
        id
        username
        email
        created_at
        role
      }
    }
  }
`;

export const LIST_USERS_QUERY = gql`
  query listUsers {
    users {
      id
      username
      email
      created_at
    }
  }
`;

export const REGISTER_USER_MUTATION = gql`
  mutation registerUser($input: CreateUserInput) {
    registerUser(input: $input) {
      id
      username
      email
      created_at
    }
  }
`;

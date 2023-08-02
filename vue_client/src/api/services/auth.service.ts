import {
  provideApolloClient,
  useMutation,
  useQuery,
} from '@vue/apollo-composable';
import {
  EMAIL_AVAILABILITY_QUERY,
  LOGIN_USER_QUERY,
  REFRESH_USER_TOKEN,
  USERNAME_AVAILABILITY_QUERY,
  REGISTER_USER_MUTATION,
} from '@/api/graphql/users.qm';
import { apolloClient } from '../apolloClient';

export const loginUser = (loginRequest: LoginRequest) => {
  return useQuery(LOGIN_USER_QUERY, {
    input: loginRequest,
  });
};

export const RegisterUser = (registerRequest: RegisterRequest) => {
  return useMutation(REGISTER_USER_MUTATION, {
    variables: { input: registerRequest },
  });
};

export const checkUsernameAvailability = (username: string) => {
  return new Promise<boolean>((resolve, reject) => {
    provideApolloClient(apolloClient)(() => {
      const { onError, onResult } = useQuery(USERNAME_AVAILABILITY_QUERY, {
        username: username,
      });
      onResult(({ data }) => {
        resolve(data.checkUsernameAvailability);
      });
      onError(() => {
        reject();
      });
    });
  });
};

export const checkEmailAvailability = (email: string) => {
  return new Promise<boolean>((resolve, reject) => {
    provideApolloClient(apolloClient)(() => {
      const { onError, onResult } = useQuery(EMAIL_AVAILABILITY_QUERY, {
        email: email,
      });
      onResult(({ data }) => {
        resolve(data.checkEmailAvailability);
      });
      onError(() => {
        reject();
      });
    });
  });
};

export const refreshToken = (token: string) => {
  return useQuery(REFRESH_USER_TOKEN, {
    token: token,
  });
};

interface LoginRequest {
  usernameOrEmail: string;
  password: string;
}

interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

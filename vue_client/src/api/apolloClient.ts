import useAuthenticationStore from '@/store/authStore';

import {
  ApolloClient,
  ApolloLink,
  createHttpLink,
  from,
  InMemoryCache,
  ServerParseError,
} from '@apollo/client/core';
import { onError } from '@apollo/client/link/error';

const httpLink = createHttpLink({
  uri: 'http://127.0.0.1:8080/query',
});

const cache = new InMemoryCache();

// Error Handling
const errorLink = onError(({ graphQLErrors, networkError }) => {
  if (graphQLErrors)
    graphQLErrors.map(({ message, locations, path }) =>
      console.log(
        `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
      )
    );
  if (networkError) console.log(`[Network error]: ${networkError}`);
});

const authMiddleware = new ApolloLink((operation, forward) => {
  // skip adding header if operation is
  if (operation.operationName === 'refreshToken') return forward(operation);

  // add the authorization to the headers
  const authStore = useAuthenticationStore();
  const token = authStore.accessToken;
  if (token) {
    operation.setContext(({ headers = {} }) => ({
      headers: {
        ...headers,
        authorization: token,
      },
    }));
  }
  return forward(operation);
});

const refreshTokenLink = onError(({ networkError }) => {
  const netError = networkError as ServerParseError;
  if (netError.statusCode === 401) {
    const authStore = useAuthenticationStore();
    if (authStore.getRefreshToken) authStore.doRefreshToken().then().catch();
  }
});

export const apolloClient = new ApolloClient({
  link: from([refreshTokenLink, errorLink, authMiddleware, httpLink]),
  cache: cache,
  connectToDevTools: true,
});

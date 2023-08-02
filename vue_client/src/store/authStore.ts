import { defineStore } from 'pinia';
import {
  loginUser,
  refreshToken,
  RegisterUser,
} from '@/api/services/auth.service';
import { provideApolloClient } from '@vue/apollo-composable';
import { apolloClient } from '@/api/apolloClient';
import { ApolloError } from '@apollo/client';
import { listUsers } from '@/api/services/users.service';
import router from '@/router';

const useAuthenticationStore = defineStore({
  id: 'authStore',
  state: (): AuthSatate => ({
    accessToken: null,
    refreshToken: null,
    currentUser: null,
    isLoading: false,
    errors: null,
  }),
  actions: {
    login(login: Login) {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const { onResult, onError } = loginUser(login);
          onError((err) => {
            this.isLoading = false;
            this.errors = err;
            this.accessToken = null;
            this.refreshToken = null;
            this.currentUser = null;
            reject(err);
          });
          onResult(({ data }) => {
            if (data.loginUser != null) {
              this.isLoading = false;
              this.accessToken = data.loginUser.access_token;
              this.refreshToken = data.loginUser.refresh_token;
              this.currentUser = data.loginUser.current_user;
              this.errors = null;
              resolve();
            }
          });
        });
      });
    },
    logout() {
      this.resetStore();
      router.push('/login');
    },
    register(register: Register) {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const {
            mutate: registerUser,
            onDone,
            onError,
          } = RegisterUser(register);
          registerUser();
          onError((err) => {
            this.isLoading = false;
            reject(err);
          });
          onDone(() => {
            this.isLoading = false;
            resolve();
          });
        });
      });
    },
    getAllUsers() {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          this.isLoading = true;
          const { onResult, onError } = listUsers();
          onError((err) => {
            this.isLoading = false;
            this.errors = err;
            reject(err);
          });
          onResult(({ data }) => {
            this.isLoading = false;
            if (data.listUsers != null) {
              console.log(data.listUsers.users);
              resolve();
            }
          });
        });
      });
    },
    doRefreshToken() {
      return new Promise<void>((resolve, reject) => {
        provideApolloClient(apolloClient)(() => {
          if (this.refreshToken == null) {
            this.logout();
          }
          const { onResult, onError } = refreshToken(
            this.refreshToken as string
          );
          onError(() => {
            this.logout();
            reject();
          });
          onResult(({ data }) => {
            if (data.refreshToken != null) {
              this.accessToken = data.refreshToken.access_token;
              this.refreshToken = data.refreshToken.refresh_token;
              this.currentUser = data.refreshToken.current_user;
              this.errors = null;
              resolve();
            }
          });
        });
      });
    },
    resetStore() {
      this.accessToken = null;
      this.refreshToken = null;
      this.currentUser = null;
      this.isLoading = false;
      this.errors = null;
      return;
    },
  },
  persist: {
    enabled: true,
    strategies: [
      {
        key: 'auth',
        storage: localStorage,
      },
    ],
  },
  getters: {
    getRefreshToken(state) {
      return state.refreshToken;
    },
    getRole(state) {
      return state.currentUser?.role;
    },
  },
});

export default useAuthenticationStore;

interface AuthSatate {
  accessToken: string | null;
  refreshToken: string | null;
  currentUser: User | null;
  isLoading: boolean;
  errors: ApolloError | null;
}

interface User {
  id: number;
  username: string;
  email: string;
  createdAt: Date;
  role: string;
}

interface Login {
  usernameOrEmail: string;
  password: string;
}

interface Register {
  username: string;
  email: string;
  password: string;
}

import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginForm from '@/components/Auth/LoginForm.vue';
import SignUpFrom from '@/components/Auth/SignUpForm.vue';
import CreateProductForm from '@/components/Product/CreateProductForm.vue';
import CreateCategoryForm from '@/components/Category/CreateCategoryForm.vue';
import ManageCategories from '@/components/Category/ManageCategories.vue';
import ManageProducts from '@/components/Product/ManageProducts.vue';
import useAuthenticationStore from '@/store/authStore';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/login',
    name: 'login',
    component: LoginForm,
    meta: {
      hideForAuth: true,
    },
  },

  {
    path: '/products/create',
    name: 'create_product',
    component: CreateProductForm,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '/products/:id',
    name: 'update_product',
    component: CreateProductForm,

    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '/products/create',
    name: 'create_product',
    component: CreateProductForm,
    meta: {
      requiresAuth: true,
    },
  },

  {
    path: '/products',
    name: 'list_products',
    component: ManageProducts,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '/categories/create',
    name: 'create_category',
    component: CreateCategoryForm,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '/categories/:id',
    name: 'update_category',
    component: CreateCategoryForm,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '/categories',
    name: 'list_categories',
    component: ManageCategories,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '/signup',
    name: 'signup',
    component: SignUpFrom,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const authStore = useAuthenticationStore();

  // Redirect to login if restricted
  if (
    to.matched.some((record) => record.meta.requiresAuth) &&
    authStore.currentUser === null
  ) {
    next({ name: 'login' });
  }

  // Prevent the login route if connected
  if (
    to.matched.some((record) => record.meta.hideForAuth) &&
    authStore.currentUser != null
  ) {
    next({ name: 'home' });
  }
  // TODO check user role
  next();
});
export default router;

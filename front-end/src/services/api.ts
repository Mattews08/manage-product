import { Category, Product } from "@/types";
import axios from "axios";

// Configura a instância do axios com a URL base da API
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
});

// Intercepta as requisições para incluir o token JWT, se presente
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("authToken");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default {
  // Método para login
  login(username: string, password: string) {
    return api.post("/login", { username, password });
  },

  // Método para obter todas as categorias
  getCategories() {
    return api.get("/categories");
  },

  // Método para criar um novo produto
  createProduct(productData: any) {
    return api.post("/products", productData);
  },

  // Método para criar uma nova categoria
  createCategory(categoryData: Category) {
    return api.post("/categories", categoryData);
  },

  // Método para deletar uma categoria
  deleteCategory(categoryId: number) {
    return api.delete(`/categories/${categoryId}`);
  },

  // Método para obter todos os produtos
  getProducts() {
    return api.get("/products");
  },

  // Método para obter um produto pelo ID
  getProductById(productId: number) {
    return api.get(`/products/${productId}`);
  },

  // Método para atualizar um produto pelo ID
  updateProduct(productId: number, productData: Product) {
    return api.put(`/products/${productId}`, productData);
  },

  // Método para deletar um produto pelo ID
  deleteProduct(productId: number) {
    return api.delete(`/products/${productId}`);
  },
};

export interface Product {
  id: number;
  name: string;
  description: string;
  imagePath: string;
  price: number;
  expiryDate: string;
  category: Category;
}

export interface DetailsProductProps {
  product: Product;
  productId: number;
  isModalOpen: boolean;
  onClose: () => void;
}

export interface Category {
  id: number;
  name: string;
}

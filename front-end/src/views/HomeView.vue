<template>
  <div class="container mx-auto px-4 py-8">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-4">
        <Select v-model="selectedCategory" class="w-48">
          <SelectTrigger>
            <SelectValue placeholder="Todas as categorias" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Todas as categorias</SelectItem>
            <SelectItem v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </SelectItem>
          </SelectContent>
        </Select>
        <div class="relative w-full max-w-md">
          <i class="pi pi-search absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground"></i>
          <Input
            type="text"
            placeholder="Pesquisar produtos..."
            v-model="searchTerm"
            class="pl-10 pr-4 py-2 rounded-md bg-background border border-input focus:border-primary focus:outline-none"
          />
        </div>
      </div>
      <div>
        <Button size="sm" variant="primary" @click="showRegisterForm = true">
          <i class="pi pi-plus mr-2"></i>
          Adicionar Novo Produto
        </Button>
        <Button size="sm" variant="secondary" @click="showCategoriesModal = true" class="ml-2">
          <i class="pi pi-tags mr-2"></i>
          Categorias
        </Button>
      </div>
    </div>
    <div class="border rounded-lg overflow-x-auto">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Nome</TableHead>
            <TableHead>Descrição</TableHead>
            <TableHead>Preço</TableHead>
            <TableHead>Data de Validade</TableHead>
            <TableHead>Categoria</TableHead>
            <TableHead>Ações</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-for="product in filteredProducts" :key="product.id">
            <TableCell class="font-medium">{{ product.name }}</TableCell>
            <TableCell>{{ product.description }}</TableCell>
            <TableCell>R$ {{ product.price.toFixed(2) }}</TableCell>
            <TableCell>{{ formatDate(product.expiryDate) }}</TableCell>
            <TableCell>{{ product.category.name }}</TableCell>
            <TableCell>
              <div class="flex items-center gap-2">
                <Button size="sm" variant="outline">
                  <i class="pi pi-eye mr-2"></i>
                  Visualizar
                </Button>
                <Button size="sm" variant="destructive">
                  <i class="pi pi-trash mr-2"></i>
                  Deletar
                </Button>
              </div>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
    <RegisterForm v-if="showRegisterForm" @close="closeRegisterForm" @product-registered="loadProducts" />

    <!-- Modal de Categorias -->
    <div v-if="showCategoriesModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div class="bg-white p-6 rounded-lg shadow-lg w-full max-w-2xl">
        <Button variant="ghost" size="icon" @click="closeCategoriesModal" class="ml-auto mb-4">
          <i class="pi pi-times w-4 h-4 text-muted-foreground"></i>
        </Button>
        <Categories />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, computed, onMounted } from 'vue'
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select'
import { Input } from '@/components/ui/input'
import { Table, TableHeader, TableRow, TableHead, TableBody, TableCell } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import RegisterForm from '@/components/RegisterForm.vue'
import Categories from '@/components/Categories.vue'
import api from '@/services/api'

export default {
  components: {
    Select, SelectTrigger, SelectValue, SelectContent, SelectItem,
    Input,
    Table, TableHeader, TableRow, TableHead, TableBody, TableCell,
    Button,
    RegisterForm,
    Categories,
  },
  setup() {
    const selectedCategory = ref("all")
    const searchTerm = ref("")
    const showRegisterForm = ref(false)
    const showCategoriesModal = ref(false)
    const products = ref([])
    const categories = ref([])

    const closeRegisterForm = () => {
      showRegisterForm.value = false
    }

    const closeCategoriesModal = () => {
      showCategoriesModal.value = false
    }

    const loadProducts = async () => {
      try {
        const response = await api.getProducts()
        products.value = response.data
      } catch (error) {
        console.error('Erro ao carregar produtos:', error)
      }
    }

    const loadCategories = async () => {
      try {
        const response = await api.getCategories()
        categories.value = response.data
      } catch (error) {
        console.error('Erro ao carregar categorias:', error)
      }
    }

    const filteredProducts = computed(() => {
      return products.value.filter((product) => {
        if (selectedCategory.value !== "all" && product.category !== selectedCategory.value) {
          return false
        }
        if (searchTerm.value.trim() !== "" && !product.name.toLowerCase().includes(searchTerm.value.toLowerCase())) {
          return false
        }
        return true
      })
    })

    const formatDate = (dateStr: string) => {
      if (!dateStr) {
        return "Data inválida";
      }

      const date = new Date(dateStr);

      if (isNaN(date.getTime())) {
        return "Data inválida";
      }

      return `${date.getDate().toString().padStart(2, '0')}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getFullYear()}`;
    }

    onMounted(() => {
      loadProducts()
      loadCategories()
    })

    return {
      selectedCategory,
      searchTerm,
      filteredProducts,
      formatDate,
      showRegisterForm,
      showCategoriesModal,
      closeRegisterForm,
      closeCategoriesModal,
      loadProducts,
    }
  }
}
</script>

<template>
  <div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
    <div class="bg-white p-6 rounded-lg shadow-lg w-full max-w-2xl">
      <Card>
        <CardHeader>
          <CardTitle>Cadastrar Novo Produto</CardTitle>
          <CardDescription>Preencha os campos abaixo para adicionar um novo produto.</CardDescription>
        </CardHeader>
        <CardContent>
          <form class="grid gap-6" @submit.prevent="submitProduct">
            <div class="grid gap-2">
              <Label for="name">Nome</Label>
              <Input id="name" v-model="product.name" placeholder="Digite o nome do produto" required />
            </div>
            <div class="grid gap-2">
              <Label for="description">Descrição</Label>
              <Textarea id="description" v-model="product.description" placeholder="Descreva o produto"
                class="min-h-[100px]" required />
            </div>
            <div class="grid sm:grid-cols-2 gap-2">
              <div class="grid gap-2">
                <Label for="price">Preço</Label>
                <Input id="price" v-model="product.price" type="number" placeholder="Digite o preço" required />
              </div>
              <div class="grid gap-2">
                <Label for="expiration">Data de Validade</Label>
                <Input id="expiration" v-model="product.expirationDate" type="date" required />
              </div>
            </div>
            <div class="grid gap-2">
              <Label for="image">Imagem</Label>
              <Input id="image" type="file" @change="handleFileUpload" />
            </div>
            <div class="grid gap-2">
              <Label for="category">Categoria</Label>
              <Select v-model="product.category" id="category" required>
                <SelectTrigger>
                  <SelectValue placeholder="Selecione a categoria" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem v-for="category in categories" :key="category.id" :value="category.id">
                    {{ category.name }}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </form>
        </CardContent>
        <CardFooter>
          <Button @click="submitProduct" class="ml-auto">
            Cadastrar Produto
          </Button>
          <Button @click="$emit('close')" class="ml-4" variant="secondary">
            Cancelar
          </Button>
        </CardFooter>
      </Card>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select'
import { Button } from '@/components/ui/button'
import api from '@/services/api'
import { toast } from 'vue3-toastify';

export default defineComponent({
  components: {
    Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter,
    Label, Input, Textarea, Select, SelectTrigger, SelectValue, SelectContent, SelectItem, Button,
  },
  setup(props, { emit }) {
    const product = ref({
      name: '',
      description: '',
      price: null,
      expirationDate: '',
      image: null,
      category: ''
    })

    const categories = ref([])

    onMounted(async () => {
      try {
        const response = await api.getCategories()
        categories.value = response.data
      } catch (error) {
        console.error('Erro ao carregar categorias:', error)
      }
    })

    const handleFileUpload = (event: Event) => {
      const target = event.target as HTMLInputElement
      if (target.files && target.files.length > 0) {
        product.value.image = target.files[0]
      }
    }

    const submitProduct = async () => {
      try {
        const formData = new FormData();
        formData.append('name', product.value.name);
        formData.append('description', product.value.description);
        formData.append('price', product.value.price.toString());
        formData.append('expirationDate', product.value.expirationDate);
        if (product.value.image) {
          formData.append('image', product.value.image);
        }
        formData.append('category', product.value.category);

        const response = await api.createProduct(formData);
        console.log('Produto cadastrado:', response.data);
        toast.success('Produto cadastrado com sucesso!');

        emit('product-registered');

        emit('close');
      } catch (error) {
        console.error('Erro ao cadastrar produto:', error);
        toast.error('Erro ao cadastrar produto!');
      }
    };

    return {
      product,
      categories,
      handleFileUpload,
      submitProduct,
    }
  }
})
</script>

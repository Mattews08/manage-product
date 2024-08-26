<template>
  <div v-if="isModalOpen" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
    <Card class="max-w-5xl mx-auto overflow-hidden bg-gradient-to-br from-background to-secondary/10 relative">
      <!-- Botão de Fechar (X) -->
      <Button variant="ghost" size="icon" class="absolute top-2 right-2" @click="$emit('close')">
        <i class="pi pi-times w-4 h-4"></i>
      </Button>
      <CardContent class="p-0">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8 p-8">
          <!-- Verifica se a imagem existe, caso contrário, aplica um plano de fundo cinza -->
          <div class="relative aspect-square rounded-xl overflow-hidden shadow-lg"
            :style="{ backgroundColor: !item.imagePath ? '#d1d5db' : 'transparent' }">
            <img v-if="item.imagePath" :src="item.imagePath" :alt="item.name" class="object-cover w-full h-full" />
          </div>
          <div class="flex flex-col justify-between space-y-6">
            <div>
              <Input v-if="isEditing" name="name" v-model="item.name" class="text-3xl font-bold mb-2" />
              <h2 v-else class="text-3xl font-bold mb-2 text-primary">{{ item.name }}</h2>
              <Textarea v-if="isEditing" name="description" v-model="item.description" class="w-full" rows="4" />
              <p v-else class="text-muted-foreground">{{ item.description }}</p>
            </div>
            <div class="space-y-4">
              <div class="flex items-center space-x-2 text-2xl font-semibold">
                <Input v-if="isEditing" name="price" type="number" v-model="item.price" class="w-32" />
                <span v-else>R$ {{ item.price.toFixed(2) }}</span>
              </div>
              <div class="flex items-center space-x-2">
                <i class="pi pi-calendar h-5 w-5 text-muted-foreground"></i>
                <Input v-if="isEditing" name="expirationDate" type="date" v-model="item.expiryDate" />
                <span v-else>Válido até: {{ formatDate(item.expiryDate) }}</span>
              </div>
              <div class="flex items-center space-x-2">
                <i class="pi pi-tag h-5 w-5 text-muted-foreground"></i>
                <Input v-if="isEditing" name="category" v-model="item.category.name" class="w-40" />
                <Badge v-else variant="outline" class="text-sm font-normal">{{ item.category.name }}</Badge>
              </div>
            </div>
            <div class="flex space-x-2">
              <Button v-if="isEditing" @click="handleSave" class="w-full">Salvar</Button>
              <Button v-else @click="handleEdit" class="w-full">
                <i class="pi pi-pencil w-4 h-4 mr-2"></i>Editar
              </Button>
              <Button variant="destructive" @click="handleDelete" class="w-full">
                <i class="pi pi-trash w-4 h-4 mr-2"></i>Deletar
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
</template>

<script lang="ts">
import { ref, watch } from 'vue';
import { Card, CardContent } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { toast } from 'vue3-toastify';
import api from '@/services/api';
import { formatDate } from '@/utils';
import { DetailsProductProps, Product } from '@/types';

export default {
  name: 'DetailsProduct',
  components: {
    Card, CardContent, Input, Textarea, Badge, Button
  },
  props: {
    isModalOpen: {
      type: Boolean,
      default: false
    },
    productId: {
      type: Number,
      required: true
    }
  },
  setup(props: DetailsProductProps, { emit }) {
    const item = ref<Product>({
      id: 0,
      name: '',
      description: '',
      imagePath: '',
      price: 0,
      expiryDate: '',
      category: { id: 0, name: '' }
    });
    const isEditing = ref(false);

    const loadProduct = async () => {
      try {
        const response = await api.getProductById(props.productId);
        item.value = response.data;
      } catch (error) {
        toast.error('Erro ao carregar o produto.');
      }
    };

    const handleEdit = () => {
      isEditing.value = true;
    };

    const handleSave = async () => {
      try {
        await api.updateProduct(props.productId, item.value);
        isEditing.value = false;
        emit('update');
        toast.success('Produto salvo com sucesso!');
      } catch (error) {
        toast.error('Erro ao salvar o produto.');
      }
    };

    const handleDelete = async () => {
      try {
        await api.deleteProduct(props.productId);
        toast.success('Produto deletado com sucesso!');
        emit('close');
        emit('update');
      } catch (error) {
        toast.error('Erro ao deletar o produto.');
      }
    };

    watch(() => props.productId, loadProduct, { immediate: true });

    return {
      item,
      isEditing,
      handleEdit,
      handleSave,
      handleDelete,
      formatDate,
    };
  }
};
</script>

<style scoped>
/* Adicione seus estilos personalizados aqui, se necessário */
</style>

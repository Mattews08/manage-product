<template>
  <div class="flex gap-6">
    <div class="flex flex-col gap-4 w-full max-w-md">
      <div v-if="categories.length === 0" class="text-center text-gray-500">
        Nenhuma categoria cadastrada.
      </div>
      <div v-else class="flex items-center justify-between bg-card p-4 rounded-md shadow-sm"
        v-for="category in categories" :key="category.id">
        <div class="flex items-center gap-3">
          <div class="bg-muted rounded-full p-2">
            <i class="pi pi-tag w-4 h-4 text-muted-foreground"></i>
          </div>
          <p class="font-medium">{{ category.name }}</p>
        </div>
        <Button variant="ghost" size="icon" @click="deleteCategory(category.id)">
          <i class="pi pi-times w-4 h-4 text-muted-foreground"></i>
        </Button>
      </div>
    </div>
    <div class="bg-card p-6 rounded-md shadow-sm w-full max-w-md">
      <h2 class="text-2xl font-bold mb-4">Adicionar Nova Categoria</h2>
      <form class="grid gap-4" @submit.prevent="addCategory">
        <div class="grid gap-1.5">
          <Label for="name">Nome</Label>
          <Input id="name" v-model="newCategoryName" placeholder="Digite o nome da categoria" />
        </div>
        <Button type="submit">Adicionar Categoria</Button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import api from '@/services/api';
import { toast } from 'vue3-toastify';

export default defineComponent({
  name: 'Categories',
  components: {
    Button,
    Label,
    Input,
  },
  setup() {
    const categories = ref([]);
    const newCategoryName = ref('');

    const loadCategories = async () => {
      try {
        const response = await api.getCategories();
        categories.value = response.data;
      } catch (error) {
        console.error('Erro ao carregar categorias:', error);
        toast.error('Erro ao carregar categorias!');
      }
    };

    const addCategory = async () => {
      if (newCategoryName.value.trim() !== '') {
        try {
          await api.createCategory({ name: newCategoryName.value.trim() });
          newCategoryName.value = '';
          loadCategories(); // Recarregar a lista de categorias
          toast.success('Categoria adicionada com sucesso!');
        } catch (error) {
          console.error('Erro ao adicionar categoria:', error);
          toast.error('Erro ao adicionar categoria!');
        }
      } else {
        toast.warn('O nome da categoria não pode estar vazio!');
      }
    };

    const deleteCategory = async (categoryId: number) => {
      try {
        await api.deleteCategory(categoryId);
        loadCategories(); // Recarregar a lista de categorias
        toast.success('Categoria excluída com sucesso!');
      } catch (error) {
        console.error('Erro ao deletar categoria:', error);
        toast.error('Erro ao excluir categoria!');
      }
    };

    onMounted(() => {
      loadCategories();
    });

    return {
      categories,
      newCategoryName,
      addCategory,
      deleteCategory,
    };
  },
});
</script>

<style scoped></style>

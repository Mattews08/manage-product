<template>
  <div class="grid lg:grid-cols-2 w-full min-h-screen">
    <div class="bg-muted hidden lg:block">
      <!-- Aqui você pode adicionar uma imagem ou outro conteúdo -->
    </div>
    <div class="flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div class="w-full max-w-md space-y-8">
        <div>
          <h2 class="text-center text-3xl font-bold tracking-tight text-gray-900">Entre com a sua conta</h2>
        </div>
        <form class="space-y-6" @submit.prevent="handleLogin">
          <div>
            <Label for="email" class="block text-sm font-medium text-gray-700">
              Usuário
            </Label>
            <div class="mt-1">
              <Input
                id="email"
                name="email"
                type="email"
                autocomplete="email"
                required
                class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-primary focus:outline-none focus:ring-primary sm:text-sm"
                v-model="email"
              />
            </div>
          </div>
          <div>
            <Label for="password" class="block text-sm font-medium text-gray-700">
              Senha
            </Label>
            <div class="mt-1">
              <Input
                id="password"
                name="password"
                type="password"
                autocomplete="current-password"
                required
                class="block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-primary focus:outline-none focus:ring-primary sm:text-sm"
                v-model="password"
              />
            </div>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <Checkbox
                id="remember-me"
                name="remember-me"
                class="h-4 w-4 text-primary focus:ring-primary"
                v-model="rememberMe"
              />
              <Label for="remember-me" class="ml-2 block text-sm text-gray-900">
                Lembre-se
              </Label>
            </div>
            <div class="text-sm">
              <RouterLink to="#" class="font-medium text-primary hover:text-primary-dark">
                Esqueceu sua senha?
              </RouterLink>
            </div>
          </div>
          <div>
            <Button
              type="submit"
              class="w-full justify-center rounded-md border border-transparent bg-primary py-2 px-4 text-sm font-medium text-primary-foreground shadow-sm hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-primary focus:ring-offset-2"
            >
              Login
            </Button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Checkbox } from '@/components/ui/checkbox'
import { Button } from '@/components/ui/button'
import { RouterLink, useRouter } from 'vue-router'
import api from '@/services/api';

export default defineComponent({
  components: {
    Label,
    Input,
    Checkbox,
    Button,
    RouterLink
  },
  setup() {
    const email = ref('');
    const password = ref('');
    const rememberMe = ref<boolean>(false);
    const router = useRouter();

    const handleLogin = async () => {
      try {
        const response = await api.login(email.value, password.value);
        const token = response.data.token;
        localStorage.setItem('authToken', token);
        router.push('/');
      } catch (error) {
        console.error('Erro ao fazer login:', error);
      }
    };

    return {
      email,
      password,
      rememberMe,
      handleLogin,
    };
  },
});
</script>

<style scoped>
/* Adicione seus estilos aqui se necessário */
</style>
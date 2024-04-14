<template>
  <div>
    <HeaderComponent />

    <div >
      <SearchComponent />
    </div>

    <!-- Contenedor principal dividido en dos partes -->
    <div >
      <EmailTableComponent />

      <TextComponent />
    </div>
  </div>
</template>

<script setup>
import HeaderComponent from "./components/HeaderComponent.vue";
import SearchComponent from "./components/SearchComponent.vue";
import EmailTableComponent from "./components/EmailTableComponent.vue";
import TextComponent from "./components/TextComponent.vue";
</script>

<style>
/* Estilos CSS globales si es necesario */
</style>

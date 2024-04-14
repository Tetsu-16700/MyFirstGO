<template>
  <div class="container">
    <!-- Contenedor principal -->
    <div class="relative">
      <!-- Contenedor del input y el ícono de búsqueda -->

      <input
        type="text"
        placeholder="Buscar"
        class="w-96 h-12 px-6 py-2 pr-12 text-lg border border-zinc-500 rounded-2xl text-zinc-500"
      />
    </div>
  </div>
</template>

<script setup>
// Lógica del componente si es necesario
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}
input {
  border-radius: 25px;
  width: calc(100% - 60px);
  padding: 10px 36px;
  position: relative;
}
</style>

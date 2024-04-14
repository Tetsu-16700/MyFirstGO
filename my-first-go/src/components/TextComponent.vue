<template>
  <div class="w-1/2 p-4">
    <table
      class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400 border-collapse"
    >
      <thead class="text-xs text-gray-700 uppercase dark:text-gray-400">
        <tr>
          <th scope="col" class="px-6 py-3 bg-gray-50 dark:bg-gray-800 border">
            Subject
          </th>
          <th scope="col" class="px-6 py-3 bg-gray-50 dark:bg-gray-800 border">
            Content
          </th>
        </tr>
      </thead>
      <tbody>
        <tr class="border-b border-gray-200 dark:border-gray-700">
          <td
            class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap bg-gray-50 dark:text-white dark:bg-gray-800 border"
          >
            Example Subject
          </td>
          <td class="px-6 py-4 border">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis dapibus est et dui vehicula, a pretium nisi faucibus. Vivamus maximus non ipsum vel maximus. Nulla consequat metus sed purus lobortis tristique. Nulla porttitor magna eget magna finibus condimentum. Mauris id felis tincidunt, malesuada metus eget, vulputate tellus. Sed et libero nec ante efficitur efficitur. Phasellus pharetra magna odio, eget elementum ligula dapibus ut. Vestibulum et lacinia sapien, eu ultricies libero. Sed vestibulum dignissim sapien eget feugiat.
          </td>
        </tr>
      </tbody>
    </table>
    <!-- Botón de responder -->
    <button class="mt-4 ml-auto px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 border border-blue-500">
      Responder
    </button>
  </div>
</template>

<script setup>
// Lógica del componente si es necesario
</script>

<style scoped>
td,
th {
  border: 1px solid #e5e7eb;
  padding: 0.85rem;
}
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}
</style>

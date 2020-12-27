<template>
  <div id="app">
    <Table />
    <AddContactForm />
  </div>
</template>

<script>
import Table from './components/Table.vue';
import AddContactForm from './components/AddContactForm.vue';

export default {
  components: { Table, AddContactForm },
  name:"App",
}
</script>

<style>

</style>

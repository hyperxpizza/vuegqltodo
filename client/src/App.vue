<template>
  <div id="app">
    <Table />
  </div>
</template>

<script>
import gql from 'graphql-tag';
import Table from './components/Table.vue';

/*
export const CONTACTS = gql`
  contacts{
    id
    firstName
    lastName
    email
    phone
  }
`
*/

export default {
  components: { Table },
  name:"App",
  apollo: {
    contacts: gql`
      query {
        contacts{
          id
          firstName
          lastName
          email
          phone
        }
      }  
    `
  }
}
</script>

<style>

</style>

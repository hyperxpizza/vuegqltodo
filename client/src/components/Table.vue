<template>
    <div id="table">
      <table>
        <caption>Vue GrapgQL TODO app with golang backend</caption>
          <thead>
            <tr>
              <th scope="col">First Name</th>
              <th scope="col">Last Name</th>
              <th scope="col">Email</th>
              <th scope="col">Phone</th>
              <th scope="col">Created at</th>
              <th scope="col">Updated at</th>
              <th scope="col">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="contact in contacts" :key="contact.id">
              <td scope="row" data-label="First Name">{{contact.firstName}}</td>
              <td data-label="Last Name">{{contact.lastName}}</td>
              <td data-label="Email">{{contact.email}}</td>
              <td data-label="Phone">{{contact.phone}}</td>
              <td data-label="Created at">{{contact.createdAt}}</td>
              <td data-label="Updated at">{{contact.updatedAt}}</td>
              </tr>
          </tbody>
      </table>
    </div>
</template>

<script>
import gql from 'graphql-tag';

export default {
    name: 'Table',
    apollo: {
      contacts: gql`
        query {
          contacts{
            id
            firstName
            lastName
            email
            phone
            createdAt
            updatedAt
        }
      }  
    `
  }
}
</script>

<style>
table {
    border: 1px solid #ccc;
    border-collapse: collapse;
    margin: 0;
    padding: 0;
    width: 100%;
    table-layout: fixed;
}

table caption {
    font-size: 1.5em;
    margin: .5em 0 .75em;
}
table tr {
  background-color: #f8f8f8;
  border: 1px solid #ddd;
  padding: .35em;
}

table th,
table td {
  padding: .625em;
  text-align: center;
}

table th {
  font-size: .85em;
  letter-spacing: .1em;
  text-transform: uppercase;
}

@media screen and (max-width: 600px) {
  table {
    border: 0;
  }

  table caption {
    font-size: 1.3em;
  }
  
  table thead {
    border: none;
    clip: rect(0 0 0 0);
    height: 1px;
    margin: -1px;
    overflow: hidden;
    padding: 0;
    position: absolute;
    width: 1px;
  }
  
  table tr {
    border-bottom: 3px solid #ddd;
    display: block;
    margin-bottom: .625em;
  }
  
  table td {
    border-bottom: 1px solid #ddd;
    display: block;
    font-size: .8em;
    text-align: right;
  }
  
  table td::before {
    /*
    * aria-label has no advantage, it won't be read inside a table
    content: attr(aria-label);
    */
    content: attr(data-label);
    float: left;
    font-weight: bold;
    text-transform: uppercase;
  }
  
  table td:last-child {
    border-bottom: 0;
  }
}
</style>
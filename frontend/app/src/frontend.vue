<template>
  <div id="app">
  <p>
    <label for="name">Ingrese el dominio: </label>
    <input
      id="name"
      v-model="name"
      type="text"
      name="name"
    >
  </p>
<p>
    <input
      type="submit"
      value="Enviar"
      v-on:click="buscarDominio()"
    >
  </p>

  <p>
    <input id='listarServidores'
      type="button"
      value="Listar los servidores"
      v-on:click="listarServidores()"
    >
  </p>
  <!--div><label class="label">Resultado: {{ list }}</label></div-->
  <p v-for="server in servers" :key="server.address">
          IP Address: {{ server.address }} <br>
          SSL Grade: {{server.ssl_grade}}<br>
          Country: {{server.country}}<br>
          Owner: {{server.owner}}<br>
  </p>
          <br>
        <p class="card-subtitle mb-2 text-muted">Domain Info: </p>
        <p class="card-text">Servers changed: {{ssl_change}}</p>
        <p class="card-text">SSL grade: {{ssl_grade}}</p>
        <p class="card-text">SSL previous grade: {{ssl_previous_grade}}</p>
        <p class="card-text">Logo: {{logo}}</p>
        <p class="card-text">Title: {{title}}</p>
        <p class="card-text">Is down: {{is_down}}</p>
  <pre id="json"> <code> {{ contenido }} </code></pre>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'App',

  methods: {
    listarServidores() {      
      axios({ method: 'GET', url: 'http://localhost:3000/listedServers'}).then(result => {
        console.log(JSON.stringify(result.data));
        //this.json = JSON.stringify(result.data, undefined, 4);
        this.contenido = JSON.stringify(result.data, undefined, "\t");
        this.$forceUpdate();
      })
    },

    buscarDominio() {
      axios({ method: 'GET', url: 'http://localhost:3000/infoServers/' + encodeURIComponent(this.name) }).then(result => {
        console.log(JSON.stringify(result.data));
        //this.contenido = JSON.stringify(result.data, undefined, "\t");
        this.servers = result.data.servers;
        this.ssl_change = result.data.server_changed;
        this.ssl_grade = result.data.ssl_grade;
        this.ssl_previous_grade = result.data.previous_ssl_grade;
        this.logo = result.data.logo;
        this.title = result.data.title;
        this.is_down = result.data.is_down;
        this.$forceUpdate();
      })
    }

  }
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

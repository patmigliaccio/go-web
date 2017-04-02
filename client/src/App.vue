<template>
  <div id="app">
    <h2>Go Web Application.</h2>
    <router-view></router-view>
    <div class="footer">
      Created by <a href="https://github.com/patmigliaccio">@patmigliaccio</a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
a {
	font-size: 1em;
	font-weight: 600;
	color: #0072c6;
	text-decoration: none;
}

a:hover {
	color: #999;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

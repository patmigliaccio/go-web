<template>
  <div class="home">
    <p>{{ msg }}</p>
    <input type="text" v-model="id" />
    <button @click="getTune">GET Tune</button>
    <div v-html="tune"></div>
    </div>
</template>

<script>
export default {
  name: 'home',
  data: function () {
    return {
      msg: 'A work in progress.',
      id: '840f77d0-4f12-49c7-a803-551641bcda5d'
    }
  },
  methods: {
    getTune: function () {
      alert(this.id)
    }
  }
}
</script>

<style scoped>

</style>

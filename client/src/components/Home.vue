<template>
	<div class="home">
		<button @click="setId">SET Tune ID</button>
		<input type="text" v-model="id" />
		<button @click="getTune">GET Tune</button>
		<div :id="'tune-' + id" v-if="tune">
			<h3 v-html="tune.Title"></h3>
			<h4 v-html="tune.Artist"></h4>
		</div>
    </div>
</template>

<script>

export default {
	name: 'home',
	data: function () {
		return {
			defaultId: '840f77d0-4f12-49c7-a803-551641bcda5d',
			id: '',
			tune: {}
		}
	},
	methods: {
		setId: function () {
			this.id = this.defaultId
		},
		getTune: function () {
			let vm = this

			this.$http.get('tune/' + vm.id)
			.then(data => {
				vm.tune = JSON.parse(data.body)
			}, err => { console.log(err) })
		}
	}
}

</script>

<style scoped>

</style>

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
			defaultId: '6a6cd781-eb18-4413-88a2-31db43da225a',
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

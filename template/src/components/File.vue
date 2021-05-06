<template>
	<div id="tag">
		<div id="icon"><img src=""></div>
		<div id="info">
			<div id="name">{{name}}</div>
			<div id="path">{{path}}</div>
			<div id="size">{{size}}</div>
		</div>
	</div>
</template>

<script lang="ts">
export default {
	name: "File",
	data: function() {
		return {
			name: "",
			path: "",
			size: ""
		}
	}
}
</script>

<style lang="less" scoped>

</style>
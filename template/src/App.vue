<template>
	<div id="pad">
		<FileSelect/>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import FileSelect from "./components/FileSelect.vue"

export default defineComponent({
	name: 'App',
	components: {
		FileSelect
	}
})
</script>

<style lang="less">
// Main frame size and structure.
body {
	width: 100%;
	height: 100%;
	margin: 0px;
	overflow: hidden;

	// Hold fullscreen and align vertical and horizontal center.
	#app {
		position: absolute;
		width: 100vw;
		height: 100vh;

		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		#pad {
			width: 300px;
			height: 450px;
			border-radius: 15px;

			background-color: #eceadf;

			transition: all 475ms;
		}

		#pad:hover {
			box-shadow: 0px 0px 15px #818181;
			background-color: #fff;
			transition: all 475ms;
		}
	}
}
</style>

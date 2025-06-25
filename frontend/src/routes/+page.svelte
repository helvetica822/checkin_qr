<script lang="ts">
	import CameraView from '$lib/components/CameraView.svelte';
	import QrDisplay from '$lib/components/QrDisplay.svelte';
	import ResultDisplay from '$lib/components/ResultDisplay.svelte';
</script>

<svelte:head>
	<title>QRコード読み取りアプリ</title>
	<meta name="description" content="QRコードをカメラで読み取るWebアプリケーション" />
</svelte:head>

<div class="app-container">
	<header class="app-header">
		<h1>QRコード読み取りアプリ</h1>
	</header>

	<div class="main-content">
		<div class="camera-section">
			<CameraView />
		</div>
		
		<div class="display-section">
			<QrDisplay />
		</div>
	</div>

	<div class="result-section">
		<ResultDisplay />
	</div>
</div>

<style>
	.app-container {
		width: 100%;
		height: 100vh;
		display: flex;
		flex-direction: column;
		background-color: #fafafa;
	}

	.app-header {
		background-color: #1976d2;
		color: white;
		padding: 16px 24px;
		text-align: center;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	}

	.app-header h1 {
		margin: 0;
		font-size: 24px;
		font-weight: 500;
	}

	.main-content {
		flex: 2;
		display: flex;
		gap: 16px;
		padding: 16px;
		min-height: 0;
	}

	.camera-section {
		flex: 1;
		min-height: 250px;
		max-height: 300px;
	}

	.display-section {
		flex: 1;
		min-height: 250px;
		max-height: 300px;
	}

	.result-section {
		flex: 1;
		padding: 16px;
		min-height: 200px;
	}

	@media (max-width: 768px) {
		.main-content {
			flex-direction: column;
			gap: 12px;
			padding: 12px;
		}

		.camera-section,
		.display-section {
			min-height: 200px;
			max-height: 250px;
		}

		.result-section {
			min-height: 150px;
		}

		.app-header h1 {
			font-size: 20px;
		}
	}
</style>

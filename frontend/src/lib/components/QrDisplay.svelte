<script lang="ts">
	import { qrImage } from '../../stores/qrStore';
	import { fade, fly } from 'svelte/transition';

	$: hasImage = $qrImage !== null;
</script>

<div class="qr-display-container">
	{#if hasImage}
		<div class="qr-image-wrapper" in:fly={{ x: -200, duration: 800 }} out:fade={{ duration: 300 }}>
			<img src={$qrImage} alt="検知されたQRコード" class="qr-image" />
			<div class="copy-animation"></div>
		</div>
	{:else}
		<div class="placeholder">
			<div class="placeholder-icon">📱</div>
			<p>QRコードをカメラに向けてください</p>
		</div>
	{/if}
</div>

<style>
	.qr-display-container {
		width: 100%;
		height: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: #f5f5f5;
		border-radius: 8px;
		position: relative;
		overflow: hidden;
	}

	.qr-image-wrapper {
		position: relative;
		max-width: 90%;
		max-height: 90%;
	}

	.qr-image {
		width: 100%;
		height: auto;
		border-radius: 8px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
		animation: pulse 2s infinite;
	}

	.copy-animation {
		position: absolute;
		top: -10px;
		left: -10px;
		right: -10px;
		bottom: -10px;
		border: 3px solid #4caf50;
		border-radius: 12px;
		animation: copyEffect 0.8s ease-out;
	}

	.placeholder {
		text-align: center;
		color: #666;
	}

	.placeholder-icon {
		font-size: 48px;
		margin-bottom: 16px;
	}

	.placeholder p {
		margin: 0;
		font-size: 16px;
	}

	@keyframes pulse {
		0%, 100% {
			transform: scale(1);
		}
		50% {
			transform: scale(1.05);
		}
	}

	@keyframes copyEffect {
		0% {
			opacity: 0;
			transform: scale(0.8);
		}
		50% {
			opacity: 1;
			transform: scale(1.1);
		}
		100% {
			opacity: 0;
			transform: scale(1);
		}
	}
</style>

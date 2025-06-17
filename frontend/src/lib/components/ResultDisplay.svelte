<script lang="ts">
	import { qrResult } from '../../stores/qrStore';
	import { fade, slide } from 'svelte/transition';

	$: hasResult = $qrResult !== null;
</script>

<div class="result-container">
	{#if hasResult}
		<div class="result-content" in:slide={{ duration: 500 }} out:fade={{ duration: 300 }}>
			<h3>読み取り結果</h3>
			<div class="result-text">
				<p>{$qrResult}</p>
			</div>
			<div class="countdown-indicator">
				<div class="countdown-bar"></div>
			</div>
		</div>
	{:else}
		<div class="placeholder">
			<p>QRコードの読み取り結果がここに表示されます</p>
		</div>
	{/if}
</div>

<style>
	.result-container {
		width: 100%;
		padding: 20px;
		background-color: #fff;
		border-radius: 8px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		min-height: 120px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.result-content {
		width: 100%;
		text-align: center;
	}

	.result-content h3 {
		margin: 0 0 16px 0;
		color: #333;
		font-size: 18px;
		font-weight: 600;
	}

	.result-text {
		background-color: #f8f9fa;
		border: 1px solid #e9ecef;
		border-radius: 4px;
		padding: 12px;
		margin-bottom: 16px;
		word-break: break-all;
	}

	.result-text p {
		margin: 0;
		font-family: 'Courier New', monospace;
		font-size: 14px;
		color: #495057;
	}

	.countdown-indicator {
		width: 100%;
		height: 4px;
		background-color: #e9ecef;
		border-radius: 2px;
		overflow: hidden;
	}

	.countdown-bar {
		height: 100%;
		background-color: #4caf50;
		animation: countdown 5s linear;
	}

	.placeholder {
		color: #666;
		text-align: center;
	}

	.placeholder p {
		margin: 0;
		font-size: 16px;
	}

	@keyframes countdown {
		from {
			width: 100%;
		}
		to {
			width: 0%;
		}
	}
</style>

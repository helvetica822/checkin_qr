<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { QRScanner } from '../../utils/qrScanner';
	import type { QRScanResult } from '../../utils/qrScanner';
	import { detectedQrCode, qrResult, qrImage, isScanning } from '../../stores/qrStore';

	let video: HTMLVideoElement;
	let canvas: HTMLCanvasElement;
	let scanner: QRScanner | null = null;
	let error: string | null = null;

	/**
	 * QRコード検証API呼び出し
	 * @param qrData - QRコードデータ
	 */
	async function verifyQRCode(qrData: string): Promise<{ valid: boolean; message: string }> {
		try {
			const response = await fetch('http://localhost:8080/api/qr-code/verify', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({ qr_data: qrData }),
			});

			if (!response.ok) {
				throw new Error('API呼び出しに失敗しました');
			}

			return await response.json();
		} catch (err) {
			console.error('QRコード検証エラー:', err);
			return { valid: false, message: '検証中にエラーが発生しました' };
		}
	}

	/**
	 * QRコード検知時の処理
	 * @param result - QRスキャン結果
	 */
	async function handleQRDetected(result: QRScanResult): Promise<void> {
		const canvas = document.createElement('canvas');
		const ctx = canvas.getContext('2d')!;
		canvas.width = result.imageData.width;
		canvas.height = result.imageData.height;
		ctx.putImageData(result.imageData, 0, 0);
		
		const imageDataUrl = canvas.toDataURL('image/png');
		
		qrImage.set(imageDataUrl);
		
		const verificationResult = await verifyQRCode(result.data);
		
		if (verificationResult.valid) {
			qrResult.set(verificationResult.message);
			detectedQrCode.set(verificationResult.message);
		} else {
			qrResult.set(verificationResult.message);
			detectedQrCode.set(verificationResult.message);
		}
		
		setTimeout(() => {
			qrImage.set(null);
			qrResult.set(null);
			detectedQrCode.set(null);
		}, 5000);
	}

	/**
	 * カメラスキャンを開始します
	 */
	async function startScanning(): Promise<void> {
		try {
			error = null;
			scanner = new QRScanner(video, canvas, handleQRDetected);
			await scanner.start();
			isScanning.set(true);
		} catch (err) {
			error = err instanceof Error ? err.message : 'カメラの開始に失敗しました';
			console.error('スキャン開始エラー:', err);
		}
	}

	/**
	 * カメラスキャンを停止します
	 */
	function stopScanning(): void {
		if (scanner) {
			scanner.stop();
			scanner = null;
		}
		isScanning.set(false);
	}

	onMount(() => {
		startScanning();
	});

	onDestroy(() => {
		stopScanning();
	});
</script>

<div class="camera-container">
	<div class="camera-view">
		<video bind:this={video} autoplay muted playsinline></video>
		<canvas bind:this={canvas} style="display: none;"></canvas>
	</div>
	
	{#if error}
		<div class="error-message">
			<p>{error}</p>
			<button on:click={startScanning}>再試行</button>
		</div>
	{/if}
</div>

<style>
	.camera-container {
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		background-color: #000;
		border-radius: 8px;
		overflow: hidden;
	}

	.camera-view {
		width: 100%;
		height: 100%;
		position: relative;
	}

	video {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.error-message {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		background-color: rgba(255, 255, 255, 0.9);
		padding: 20px;
		border-radius: 8px;
		text-align: center;
	}

	.error-message p {
		margin: 0 0 10px 0;
		color: #d32f2f;
	}

	.error-message button {
		background-color: #1976d2;
		color: white;
		border: none;
		padding: 8px 16px;
		border-radius: 4px;
		cursor: pointer;
	}

	.error-message button:hover {
		background-color: #1565c0;
	}
</style>

<script lang="ts">
	import { qrResult } from '../../stores/qrStore';
	import { fade, slide } from 'svelte/transition';

	$: hasResult = $qrResult !== null;
	
	// QRコード結果からユーザIDを抽出する関数
	function extractUserIdFromResult(result: string | null): string | null {
		if (!result) return null;
		
		// "不正なQRコードです" の場合はそのまま返す
		if (result === '不正なQRコードです') {
			return result;
		}
		
		// QRコードの形式は "user_id:random_string" なので、:で分割してuser_idを取得
		const parts = result.split(':');
		if (parts.length >= 2) {
			return parts[0]; // user_id部分を返す
		}
		
		// 形式が異なる場合はそのまま返す
		return result;
	}
	
	$: userId = extractUserIdFromResult($qrResult);
	$: welcomeMessage = userId && userId !== '不正なQRコードです' ? `${userId}ようこそ` : userId;
</script>

<div class="result-container">
	{#if hasResult}
		<div class="result-content" in:slide={{ duration: 500 }} out:fade={{ duration: 300 }}>
			<div class="welcome-message">
				<p class="welcome-text">{welcomeMessage}</p>
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
		height: 100%;
		padding: 20px;
		background-color: #fff;
		border-radius: 8px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.result-content {
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		text-align: center;
	}

	.welcome-message {
		flex: 1;
		display: flex;
		align-items: center;
		justify-content: center;
		margin-bottom: 20px;
	}

	.welcome-text {
		margin: 0;
		font-size: 2.5rem;
		font-weight: 600;
		color: #1976d2;
		text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.1);
		line-height: 1.2;
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

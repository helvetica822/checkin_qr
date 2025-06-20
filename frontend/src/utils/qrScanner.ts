import jsQR from 'jsqr';

export interface QRScanResult {
	data: string;
	imageData: ImageData;
}

/**
 * カメラからQRコードをスキャンするクラス
 */
export class QRScanner {
	private video: HTMLVideoElement;
	private canvas: HTMLCanvasElement;
	private context: CanvasRenderingContext2D;
	private stream: MediaStream | null = null;
	private animationId: number | null = null;
	private onDetected: (result: QRScanResult) => void;

	constructor(
		video: HTMLVideoElement,
		canvas: HTMLCanvasElement,
		onDetected: (result: QRScanResult) => void
	) {
		this.video = video;
		this.canvas = canvas;
		this.context = canvas.getContext('2d')!;
		this.onDetected = onDetected;
	}

	/**
	 * カメラを開始してQRコードスキャンを開始します
	 * @throws カメラアクセスに失敗した場合
	 */
	async start(): Promise<void> {
		try {
			this.stream = await navigator.mediaDevices.getUserMedia({
				video: { facingMode: 'environment' }
			});
			this.video.srcObject = this.stream;
			this.video.play();
			this.video.addEventListener('loadedmetadata', () => {
				this.canvas.width = this.video.videoWidth;
				this.canvas.height = this.video.videoHeight;
				this.scan();
			});
		} catch (error) {
			console.error('カメラアクセスに失敗しました:', error);
			throw new Error('カメラにアクセスできませんでした');
		}
	}

	/**
	 * QRコードスキャンを停止します
	 */
	stop(): void {
		if (this.animationId) {
			cancelAnimationFrame(this.animationId);
			this.animationId = null;
		}
		if (this.stream) {
			this.stream.getTracks().forEach(track => track.stop());
			this.stream = null;
		}
	}

	private scan(): void {
		if (this.video.readyState === this.video.HAVE_ENOUGH_DATA) {
			this.context.drawImage(this.video, 0, 0, this.canvas.width, this.canvas.height);
			const imageData = this.context.getImageData(0, 0, this.canvas.width, this.canvas.height);
			const code = jsQR(imageData.data, imageData.width, imageData.height);

			if (code) {
				this.onDetected({
					data: code.data,
					imageData: imageData
				});
			}
		}
		this.animationId = requestAnimationFrame(() => this.scan());
	}
}

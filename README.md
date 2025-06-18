# QRコード生成API & QRコード読み取りWebアプリケーション

ユーザIDを受け取ってQRコードを生成するAPI（Go + Echo）と、カメラでQRコードを読み取るWebアプリケーション（TypeScript + Svelte）のプロジェクトです。

## 技術スタック

### バックエンド
- **言語**: Go 1.21
- **Webフレームワーク**: Echo v4.11.4
- **データベース**: SQLite（インメモリ）
- **QRコード生成**: github.com/skip2/go-qrcode
- **その他**: CORS対応、ログ出力、リカバリー機能

### フロントエンド
- **言語**: TypeScript
- **フレームワーク**: Svelte 4.0.5 + SvelteKit 1.20.4
- **ビルドツール**: Vite 4.4.2
- **UIライブラリ**: Svelte Material UI (@smui/*)
- **QRコード読み取り**: jsQR 1.4.0
- **その他**: WebRTC（カメラアクセス）、レスポンシブデザイン

## 処理の仕様

### API仕様

#### QRコード生成API
- **エンドポイント**: `POST /api/qr-code/generate`
- **リクエスト**:
  ```json
  {
    "user_id": "ユーザID（文字列）"
  }
  ```
- **レスポンス**: PNG形式のQRコード画像
- **処理内容**:
  1. ユーザIDを受け取る
  2. 10文字のランダム文字列を生成
  3. "user_id:random_string" 形式でQRコードを作成
  4. データベースに保存（同一ユーザIDは上書き）
  5. PNG画像として返却

#### ヘルスチェックAPI
- **エンドポイント**: `GET /api/health`
- **レスポンス**: サーバー稼働状況

### UI仕様

#### レイアウト
- **左側（50%）**: カメラビュー
- **右側（50%）**: QRコード表示エリア
- **下部**: 読み取り結果表示エリア

#### 機能
- **カメラアクセス**: WebRTC APIを使用したリアルタイムカメラ映像
- **QRコード検知**: jsQRライブラリによる自動検知
- **アニメーション**: カメラからQRコード表示エリアへのコピー効果
- **自動消去**: 検知から5秒後に画像と結果を自動削除
- **エラーハンドリング**: カメラアクセス失敗時の再試行機能

### データベース仕様

#### qr_codesテーブル
```sql
CREATE TABLE qr_codes (
    user_id TEXT PRIMARY KEY,
    random_string TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
```

## 実行方法

### 前提条件
- Go 1.21以上
- Node.js 18以上
- npm または yarn

### バックエンドの起動

1. バックエンドディレクトリに移動
   ```bash
   cd backend
   ```

2. 依存関係のインストール
   ```bash
   go mod tidy
   ```

3. サーバー起動
   ```bash
   go run main.go
   ```

4. 動作確認
   ```bash
   # ヘルスチェック
   curl http://localhost:8080/api/health
   
   # QRコード生成テスト
   curl -X POST http://localhost:8080/api/qr-code/generate \
     -H "Content-Type: application/json" \
     -d '{"user_id": "test_user"}' \
     --output test_qr.png
   ```

### フロントエンドの起動

1. フロントエンドディレクトリに移動
   ```bash
   cd frontend
   ```

2. 依存関係のインストール
   ```bash
   npm install
   ```

3. 開発サーバー起動
   ```bash
   npm run dev
   ```

4. ブラウザでアクセス
   ```
   http://localhost:5173
   ```

### 本番ビルド

#### フロントエンド
```bash
cd frontend
npm run build
npm run preview  # ビルド結果の確認
```

#### バックエンド
```bash
cd backend
go build -o qr-backend main.go
./qr-backend
```

## 開発情報

### ディレクトリ構成
```
qr-code-app/
├── backend/
│   ├── main.go              # メインサーバー
│   ├── go.mod               # Go依存関係
│   ├── database/
│   │   └── db.go           # データベース接続
│   ├── handlers/
│   │   └── qr_handler.go   # APIハンドラー
│   ├── models/
│   │   └── qr_code.go      # データモデル
│   └── utils/
│       └── random.go       # ユーティリティ関数
└── frontend/
    ├── package.json         # npm依存関係
    ├── vite.config.js       # Vite設定
    ├── svelte.config.js     # Svelte設定
    ├── tsconfig.json        # TypeScript設定
    └── src/
        ├── routes/          # ページルーティング
        ├── lib/components/  # Svelteコンポーネント
        ├── stores/          # 状態管理
        └── utils/           # ユーティリティ関数
```

### 注意事項
- データベースはインメモリSQLiteを使用しているため、サーバー再起動でデータが消失します
- カメラアクセスはHTTPS環境またはlocalhostでのみ動作します
- QRコード読み取り精度は照明条件に依存します

### 今後の改善予定
- PostgreSQLへのデータベース移行
- Docker対応
- テストコードの追加
- CI/CD パイプラインの構築

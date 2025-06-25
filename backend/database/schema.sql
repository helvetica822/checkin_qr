CREATE TABLE IF NOT EXISTS qr_codes (
    user_id VARCHAR(255) PRIMARY KEY,
    random_string VARCHAR(10) NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_qr_codes_user_id_random_string 
ON qr_codes(user_id, random_string);

COMMENT ON TABLE qr_codes IS 'QRコード生成・検証用テーブル';
COMMENT ON COLUMN qr_codes.user_id IS 'ユーザーの一意識別子';
COMMENT ON COLUMN qr_codes.random_string IS 'QRコード用ランダム文字列（10文字）';
COMMENT ON COLUMN qr_codes.status IS 'QRコード読み取り状態（0:未読み取り、1:読み取り済み）';
COMMENT ON COLUMN qr_codes.created_at IS 'レコード作成日時';
COMMENT ON COLUMN qr_codes.updated_at IS 'レコード更新日時';

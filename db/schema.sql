CREATE TABLE IF NOT EXISTS todos (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  subject TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  created_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  CHECK(subject <> '') -- check成約でデータを追加する時に値が指定した条件を満たしているかどうかのチェック
  -- この場合は､subject列が空でないことを確認する
);

CREATE TRIGGER IF NOT EXISTS trigger_todos_updated_at
AFTER
UPDATE
  ON todos BEGIN
UPDATE
  todos
SET
  updated_at = DATETIME('now')
WHERE
  id == NEW.id;

END;
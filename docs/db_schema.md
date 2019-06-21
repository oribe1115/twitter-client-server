# db スキーマ

## stamps

| カラム名         | 型             | 説明など                             |
| :--------------- | :------------- | :----------------------------------- |
| stamp_id         | string         | twemoji の絵文字の id                |
| tweet_id         | integer(int64) | ツイートの id                        |
| user_id          | integer(int64) | スタンプを押した人のユーザー id      |
| user_screen_name | string         | スタンプを押した人のユーザー名(@...) |
| count            | integer        | 押されたスタンプの数                 |
| created_at       | string         | 最初に押した日時                     |
| updated_at       | string         | 最後に押した日時                     |

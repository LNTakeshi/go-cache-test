# goのテストをキャッシュする

mainブランチもキャッシュして、PR作った初回からキャッシュ有効にしちゃうぞ

5秒間待つだけのテストが、キャッシュによって0秒で終わることの確認ができます


## test.yml

普通にキャッシュを使うが、setup-goのキャッシュを無効にしておくことが重要

lint等もlint用のキャッシュを作っておくこと


## cache.yml

mainブランチにpushされたことをトリガーに、キャッシュを作成する

go.modが定期的に更新されることが前提になっているので、プロジェクトの状況によっては週1で必ずやる、とかでも良いかもしれない

package main

import (
	"net/http"
)

// TODO は今回のデータペイロードである
type TODO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	User string `json:"user"`
	Done bool   `json:"done"`
}

var (
// TODO mapを並行でアクセスできるようにするためのLockオブジェクトを生成
// TODO TODOstructをintをキーにして引けるようにmapを生成
// TODO IDを発番するためにint変数を生成
)

func Get(w http.ResponseWriter, r *http.Request) {
	// TODO リクエストURIからIDを取得する    ヒント r.RequestURIを参照

	// TODO IDに応じたデータをmapから取得

	if false { // TODO mapになかった場合はhttp.StatusNotFoundを返却
	}

	// TODO mapにあった場合はhttp.StatusOKを返却

	// TODO mapから取得した値をjsonに変換しResponseWriterに書き込む
}

func Create(w http.ResponseWriter, r *http.Request) {
	// TODO リクエストBodyをjson Decodeする

	if false { // TODO Decodeに失敗した場合はStatusBadRequestを返却
	}

	if false { // TODO todo.IDを確認し0であれば現在の最大値をセットし最大値を更新する
	}

	// TODO mapにtodoをtodo.IDをキーにして格納

	// TODO StatusAcceptedを返却

	// TODO todoをjsonに変換しResponseWriterに書き込む
}

func Update(w http.ResponseWriter, r *http.Request) {
	// TODO リクエストURIからIDを取得する    ヒント r.RequestURIを参照

	// TODO IDに応じたデータをmapから取得

	if false { // TODO mapになかった場合はhttp.StatusNotFoundを返却
	}

	// TODO todo.Doneをtrueにする

	// TODO mapにtodoをtodo.IDをキーにして格納

	// TODO StatusAcceptedを返却

	// TODO todoをjsonに変換しResponseWriterに書き込む

}

func Delete(w http.ResponseWriter, r *http.Request) {
	// TODO リクエストURIからIDを取得する    ヒント r.RequestURIを参照

	// TODO mapからtodoをtodo.IDをキーにして削除

	// TODO StatusAcceptedを返却
}

func main() {

	var port string
	// TODO 環境変数PORTから起動portを取得

	mux := http.DefaultServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO リクエストメソッド (r.Method) に応じて、適切なHandler (Get, Create, Update, Delete) にルーティング
		// TODO GET POST PUT PATCH DELETE 以外のリクエストはMethodNotAllowedを返却
	})

	http.ListenAndServe(port, mux)
}

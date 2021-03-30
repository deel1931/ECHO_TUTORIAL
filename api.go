package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

//User is param.
type User struct {
	ID      int    `json:"id"`
	GroupID int    `json:"group_id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
}

//echoの構造体
// Echo struct {
// 	common
// 	// startupMutex is mutex to lock Echo instance access during server configuration and startup. Useful for to get
// 	// listener address info (on which interface/port was listener binded) without having data races.
// 	startupMutex     sync.RWMutex
// 	StdLogger        *stdLog.Logger
// 	colorer          *color.Color
// 	premiddleware    []MiddlewareFunc
// 	middleware       []MiddlewareFunc
// 	maxParam         *int
// 	router           *Router
// 	routers          map[string]*Router
// 	notFoundHandler  HandlerFunc
// 	pool             sync.Pool
// 	Server           *http.Server
// 	TLSServer        *http.Server
// 	Listener         net.Listener
// 	TLSListener      net.Listener
// 	AutoTLSManager   autocert.Manager
// 	DisableHTTP2     bool
// 	Debug            bool
// 	HideBanner       bool
// 	HidePort         bool
// 	HTTPErrorHandler HTTPErrorHandler
// 	Binder           Binder
// 	Validator        Validator
// 	Renderer         Renderer
// 	Logger           Logger
// 	IPExtractor      IPExtractor
// 	ListenerNetwork  string
// }

func main() {
	e := echo.New()
	// インスタンス化
	// 	func New() (e *Echo) {
	// 	e = &Echo{
	// 		Server:    new(http.Server), 　　   newは指定した型のポインタ型を生成する  構造体の初期化でよく使う
	// 		TLSServer: new(http.Server),        http.Serverはカスタムサーバー
	// 		AutoTLSManager: autocert.Manager{   ACMEプロトコルをサポート 　公開鍵証明書の標準プロトコル
	// 			Prompt: autocert.AcceptTOS,    アカウント登録時に認証局の利用規約に同意したことを示すために常にtrueを返す。
	// 		},
	// 		Logger:          log.New("echo"),  新しいlogの作成
	// 		colorer:         color.New(),
	// 		maxParam:        new(int),        許可される最大パスパラメータを設定 デフォルト５
	// 		ListenerNetwork: "tcp",
	// 	}
	// 	e.Server.Handler = e
	// 	e.TLSServer.Handler = e
	// 	e.HTTPErrorHandler = e.DefaultHTTPErrorHandler   デフォルトのHTTPエラーハンドラー  137
	// 	e.Binder = &DefaultBinder{}      バインダーは、Bindメソッドをラップするインターフェース  153
	// 	e.Logger.SetLevel(log.ERROR)     セットしたレベル以上のレベルを出力
	// 	e.StdLogger = stdLog.New(e.Logger.Output(), e.Logger.Prefix()+": ", 0)  標準出力へのシンプルで高速なロギングを提供
	// 	e.pool.New = func() interface{} {
	// 		return e.NewContext(nil, nil)   CONTEXTインスタンスを返す  215
	// 	}
	// 	e.router = NewRouter(e)             routerインスタンスを返す　227
	// 	e.routers = map[string]*Router{}
	// 	return
	// }

	routing(e)

	e.Logger.Fatal(e.Start(":1313"))
}

func routing(e *echo.Echo) {
	e.GET("/", hello)
	e.GET("/api/v1/groups/:group_id/users", sendjson)
	// 	func (e *Echo) GET(path string, h HandlerFunc, m ...MiddlewareFunc) *Route
	// {
	// 		return e.Add(http.MethodGet, path, h, m...)
	// 	}

	// Route struct {
	// 	Method string `json:"method"`
	// 	Path   string `json:"path"`
	// 	Name   string `json:"name"`
	// }

	// func (e *Echo) Add(method, path string, handler HandlerFunc, middleware ...MiddlewareFunc) *Route {
	// 	return e.add("", method, path, handler, middleware...)
	// }

}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	//statusOKはhttpの200を表し、リクエストが成功したことを表す。
}

func sendjson(c echo.Context) error {
	//contextは現在のHTTPリクエスト状況
	groupIDStr := c.Param("group_id")
	//Pramはパスパラメータを意味する
	groupID, err := strconv.Atoi(groupIDStr)
	//strconvは基本的な方の文字列表現への変換関数
	//Atoiは文字列を整数にする
	if err != nil {
		return xerrors.Errorf("errors when group id convert to int %s: %w", groupIDStr, err)
	}
	//Wrpfは直前の関数やメッドの履歴に対してのエラーに注釈を付けられる。
	//ラッパーとは元の関数は変えずに窓口役として立ってくれる
	gender := c.QueryParam("gender")
	//URLの？についてるパラメータの読み込み
	users := []*User{} //map
	if gender == "" || gender == "man" {
		users = append(users, &User{ID: 1, GroupID: groupID, Name: "Taro", Gender: "man"})
		users = append(users, &User{ID: 2, GroupID: groupID, Name: "Jiro", Gender: "man"})
	}
	if gender == "" || gender == "woman" {
		users = append(users, &User{ID: 3, GroupID: groupID, Name: "Hanako", Gender: "woman"})
		users = append(users, &User{ID: 4, GroupID: groupID, Name: "Yoshiko", Gender: "woman"})
	}
	return c.JSON(http.StatusOK, users)
	//json形式で返す。
}

//DefaultHTTPErrorHandlerは、デフォルトのHTTPエラーハンドラーです。 JSON応答を送信します
//ステータスコード付き
// func (e *Echo) DefaultHTTPErrorHandler(err error, c Context) {
// 	he, ok := err.(*HTTPError)
// 	if ok {
// 		if he.Internal != nil {
// 			if herr, ok := he.Internal.(*HTTPError); ok {
// 				he = herr
// 			}
// 		}
// 	} else {
// 		he = &HTTPError{
// 			Code:    http.StatusInternalServerError,  //webサーバ上でエラーが出て画面表示できない状態
// 			Message: http.StatusText(http.StatusInternalServerError),
// 		}
// 	}

//binder methods
// func (b *DefaultBinder) Bind(i interface{}, c Context) (err error) {
// 	if err := b.BindPathParams(c, i); err != nil {
// 		return err
// 	}

// パスパラメータをバインド可能なオブジェクトにバインドします
// func (b *DefaultBinder) BindPathParams(c Context, i interface{}) error {
// 	names := c.ParamNames()  311
// 	values := c.ParamValues() 332
// 	params := map[string][]string{}
// 	for i, name := range names {
// 		params[name] = []string{values[i]}
// 	}
// 	if err := b.bindData(i, params, "param"); err != nil {
// 		return NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
// 	}
// 	return nil
// }

//pnamesを返す
// func (c *context) ParamNames() []string {
// 	return c.pnames　　構造体の中に定義されている　文字列型
// }

//pvaluesのなかのpnamesの数を返す
// func (c *context) ParamValues() []string {
// 	return c.pvalues[:len(c.pnames)]
// }

//bindDataは、EXPLICITタグを持つ宛先構造体のデータのみのフィールドをバインドします　翻訳
// func (b *DefaultBinder) bindData(destination interface{}, data map[string][]string, tag string) error {
// 	if destination == nil || len(data) == 0 {
// 		return nil
// 	}
// 	typ := reflect.TypeOf(destination).Elem()  typeOfは型の取得
// 	val := reflect.ValueOf(destination).Elem() .Elem() でポインタの値を返している

// 	// Map
// 	if typ.Kind() == reflect.Map {
// 		for k, v := range data {
// 			val.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v[0]))
// 		}
// 		return nil
// 	}

// NewHTTPErrorは、新しいHTTPErrorインスタンスを作成します。
// func NewHTTPError(code int, message ...interface{}) *HTTPError {
// 	he := &HTTPError{Code: code, Message: http.StatusText(code)}
// 	if len(message) > 0 {
// 		he.Message = message[0]
// 	}
// 	return he
// }

// SetInternalはエラーをHTTPError.Internalに設定します
// func (he *HTTPError) SetInternal(err error) *HTTPError {
// 	he.Internal = err
// 	return he
// }

// NewContextはContextインスタンスを返します。
// func (e *Echo) NewContext(r *http.Request, w http.ResponseWriter) Context {
// 	return &context{
// 		request:  r,
// 		response: NewResponse(w, e),
// 		store:    make(Map),
// 		echo:     e,
// 		pvalues:  make([]string, *e.maxParam),
// 		handler:  NotFoundHandler,
// 	}
// }

// NewRouter は新しいルーターインスタンスを返す
// func NewRouter(e *Echo) *Router {
// 	return &Router{
// 		tree: &node{
// 			methodHandler: new(methodHandler),    238
// 		},
// 		routes: map[string]*Route{},
// 		echo:   e,
// 	}
// }

//methodhandler
// methodHandler struct {
// 	connect  HandlerFunc
// 	delete   HandlerFunc
// 	get      HandlerFunc
// 	head     HandlerFunc
// 	options  HandlerFunc
// 	patch    HandlerFunc
// 	post     HandlerFunc
// 	propfind HandlerFunc
// 	put      HandlerFunc
// 	trace    HandlerFunc
// 	report   HandlerFunc
// }

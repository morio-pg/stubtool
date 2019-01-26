package model

var (
	// ErrInitializingApp アプリの初期化エラー
	ErrInitializingApp = "error initializing app: %v"
	// ErrGettingAuthClient 認証クライアントの取得エラー
	ErrGettingAuthClient = "error getting Auth client: %v"
	// ErrVerifyingIDToken IDトークンの検証エラー
	ErrVerifyingIDToken = "error verifying ID token: %v"
	// ErrDeletingUser ユーザーの削除エラー
	ErrDeletingUser = "error deleting user: %v"
	// ErrAuthorizationHeader Authorization Headerエラー
	ErrAuthorizationHeader = "error Authorization Header: %v"
	// ErrGetData データ取得エラー
	ErrGetData = "error get data: %v"
	// ErrConvertData データ変換エラー
	ErrConvertData = "error convert data: %v"
	// ErrSaveData データ保存エラー
	ErrSaveData = "error save data: %v"
	// ErrDeleteData データ削除エラー
	ErrDeleteData = "error delete data: %v"
	// ErrMismatchUID UID不一致エラー
	ErrMismatchUID = "error mismatch uid: %s != %s"
	// ErrMismatchFilename ファイル名不一致エラー
	ErrMismatchFilename = "error mismatch filename: %s != %s"
)

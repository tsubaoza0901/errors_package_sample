package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

// --------
// model↓
// --------

// User ...
type User struct {
	ID             uint   `json:"id" param:"id"`       // paramタグ
	Name           string `json:"name" query:"name"`   // queryタグ
	Age            int    `json:"age" query:"age"`     // queryタグ
	SomethingArray []int  `json:"array" query:"array"` // queryタグ
}

// --------
// router↓
// --------

// InitRouting ...
func InitRouting(e *echo.Echo, u *User) {
	// e.GET("something", SomethingHandler)
	// e.POST("user", u.CreateUser)
	// e.PUT("user/:id", u.UpdateUser)
	// e.DELETE("user/:id", u.DeleteUser)
	e.GET("user/:id", u.GetUser)
	// e.GET("users", u.GetUsers)
}

// --------
// handler↓
// --------

// // SomethingHandler ...
// func SomethingHandler(c echo.Context) error {
// 	_, err := os.Open("xxx")
// 	if err != nil {
// 		code := ErrFailedToServer
// 		err = errors.Cause(err)
// 		fmt.Println(err)
// 		if apperr, ok := err.(*AppError); ok {
// 			code = apperr.Code
// 			err = apperr.Wrap()
// 		} else {
// 			err = NewAppError(ErrFailedToServer, err).Wrap()
// 		}
// 		return c.JSON(http.StatusOK, NewAPIResponse(code, err.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, "Done!")
// }

// // CreateUser ...
// func (u *User) CreateUser(c echo.Context) error {
// 	user := User{}

// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}

// 	user = User{
// 		ID:   1,
// 		Name: user.Name,
// 		Age:  user.Age,
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// // UpdateUser ...
// func (u *User) UpdateUser(c echo.Context) error {
// 	user := User{}

// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, "Updated")
// }

// // DeleteUser ...
// func (u *User) DeleteUser(c echo.Context) error {
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusOK, "Deleted")
// }

// GetUser ...
func (u *User) GetUser(c echo.Context) error {
	user := User{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		log.Error(err)

		// fmt.Println(NewAppError(ErrBadRequest, err))
		err = NewAppError(ErrBadRequest, err).Wrap()
		// fmt.Println(err)
		return c.JSON(http.StatusOK, NewAPIResponse(ErrBadRequest, err.Error(), nil))
	}

	// Getメソッドのイメージ
	if id == 1 {
		user = User{
			ID:   1,
			Name: "Tom",
			Age:  29,
		}
	} else if id == 2 {
		user = User{
			ID:   2,
			Name: "Bob",
			Age:  35,
		}

	} else {
		// fmt.Println(err) // => <nil>
		// log.Error(err)
		err = errors.New("Record not found")
		// fmt.Println(err) // => Record not found
		log.Error(err)
		err = NewAppError(ErrRecordNotFound, err).Wrap()
		return c.JSON(http.StatusOK, NewAPIResponse(ErrRecordNotFound, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, NewAPIResponse(0, StatusText(StatusSuccess), user))
}

// // GetUsers ...
// func (u *User) GetUsers(c echo.Context) error {
// 	users := []*User{}

// 	name := c.QueryParam("name")

// 	// Get Allメソッドのイメージ
// 	if name == "" {
// 		users = []*User{
// 			{
// 				ID:   1,
// 				Name: "Tom",
// 				Age:  29,
// 			},
// 			{
// 				ID:   2,
// 				Name: "Bob",
// 				Age:  35,
// 			},
// 		}
// 	} else if name == "Tom" {
// 		users = []*User{
// 			{
// 				ID:   1,
// 				Name: "Tom",
// 				Age:  29,
// 			},
// 		}
// 	} else if name == "Bob" {
// 		users = []*User{
// 			{
// 				ID:   2,
// 				Name: "Bob",
// 				Age:  35,
// 			},
// 		}
// 	} else {
// 		return c.JSON(http.StatusOK, "Not Found")
// 	}

// 	return c.JSON(http.StatusOK, users)
// }

// --------
// errors
// --------

// APIResponse レスポンス用構造体
type APIResponse struct {
	Status   string      `json:"status" example:"A400"`
	Message  string      `json:"msg" example:"some error"`
	Response interface{} `json:"response"`
}

// APIResponseのStatusコード定義
const (
	StatusSuccess       = "A200"
	StatusBadRequestErr = "A400"
	StatusNotFoundErr   = "A404"
	StatusServerErr     = "A500"
	StatusUnauthorized  = "A401"
)

var statusText = map[string]string{
	StatusSuccess: "Success",
}

// StatusText APIResponseのStatusを元にしてエラーメッセージ取得する関数
func StatusText(status string) string {
	return statusText[status]
}

// NewAPIResponse 処理結果を元にしてレスポンス情報生成する関数
func NewAPIResponse(errCode int, msg string, res interface{}) *APIResponse {
	sts := StatusSuccess
	switch errCode {
	case ErrBadRequest,
		ErrExistSameName,
		ErrUsedDesignTemplate,
		ErrExclusionControl,
		ErrUsedPlacementRelation,
		ErrStandardTemplate,
		ErrExistSameAdID,
		ErrUsedDeviceModel,
		ErrUsedDeviceModelGroup,
		ErrUsedOS,
		ErrUsedOSGroup,
		ErrUsedSDKVersion,
		ErrExistSameOuterDealID,
		ErrChangeOuterDealID,
		ErrUnsupportedPMP,
		ErrUsedDeal,
		ErrUnsupportedAutomaticRatio,
		ErrChangeAdFormat,
		ErrChangeDesignTemplateType,
		ErrChangeMediaType,
		ErrChangePlatform,
		ErrExistSameTemplateTypePlatformAdFormat,
		ErrNoPermission,
		ErrExistEncryptKeyPlatform:
		sts = StatusBadRequestErr
	case ErrRecordNotFound:
		sts = StatusNotFoundErr
	case ErrFailedToServer:
		sts = StatusServerErr
	case ErrUnauthorized:
		sts = StatusUnauthorized
	}

	return &APIResponse{Status: sts, Message: msg, Response: res}
}

// 出力するエラーメッセージのインデックス（Error Code）
const (
	_ = iota
	ErrBadRequest
	ErrRecordNotFound
	ErrExistSameName
	ErrUsedDesignTemplate
	ErrFailedToServer
	ErrExclusionControl
	ErrUnauthorized
	ErrUsedPlacementRelation
	ErrStandardTemplate
	ErrExistSameAdID
	ErrUsedDeviceModel
	ErrUsedDeviceModelGroup
	ErrUsedOS
	ErrUsedOSGroup
	ErrUsedSDKVersion
	ErrExistSameOuterDealID
	ErrChangeOuterDealID
	ErrUnsupportedPMP
	ErrUsedDeal
	ErrUnsupportedAutomaticRatio
	ErrChangeAdFormat
	ErrChangeDesignTemplateType
	ErrChangeMediaType
	ErrChangePlatform
	ErrExistSameTemplateTypePlatformAdFormat
	ErrNoPermission
	ErrExistEncryptKeyPlatform
)

// Wrapするエラーメッセージ内容
var errorText = map[int]string{
	ErrBadRequest:                            "不正な値が設定されています",
	ErrRecordNotFound:                        "データが削除されているか存在しません",
	ErrExistSameName:                         "同じ名前は登録できません",
	ErrUsedDesignTemplate:                    "配信セットで使用されているため、削除できません",
	ErrFailedToServer:                        "予期せぬエラーが発生しました",
	ErrExclusionControl:                      "他のユーザーにより更新されました。再度更新してやり直して下さい",
	ErrUnauthorized:                          "認証エラー",
	ErrUsedPlacementRelation:                 "配信セットで使用されているため、削除できません",
	ErrStandardTemplate:                      "標準テンプレートは、編集/削除できません",
	ErrExistSameAdID:                         "同じADID/IDFAは登録できません",
	ErrUsedDeviceModel:                       "機種グループで使用されているため、削除できません",
	ErrUsedDeviceModelGroup:                  "配信セットで使用されているため、削除できません",
	ErrUsedOS:                                "OSグループで使用されているため、削除できません",
	ErrUsedOSGroup:                           "配信セットまたはプレースメント設定で使用されているため、削除できません",
	ErrUsedSDKVersion:                        "SDKバージョングループで使用されているため、削除できません",
	ErrExistSameOuterDealID:                  "同じDealIDは登録できません",
	ErrChangeOuterDealID:                     "編集時にDealIDの変更はできません",
	ErrUnsupportedPMP:                        "PMP非対応のPFにDealの選択はできません",
	ErrUsedDeal:                              "配信セットで使用されているため、削除できません",
	ErrUnsupportedAutomaticRatio:             "配信比率自動化非対応のPFの選択はできません",
	ErrChangeAdFormat:                        "編集時にフォーマットの変更はできません",
	ErrChangeDesignTemplateType:              "編集時に配信手法の変更はできません",
	ErrChangeMediaType:                       "編集時に配信面の変更はできません",
	ErrExistEncryptKeyPlatform:               "既にkey発行済のプラットフォームです",
	ErrChangePlatform:                        "編集時にプラットフォームの変更はできません",
	ErrExistSameTemplateTypePlatformAdFormat: "同じ配信手法/プラットフォーム/フォーマットは登録できません",
	ErrNoPermission:                          "登録編集する権限がありません",
}

// ErrorText 受け取ったError Code（エラーメッセージのインデックス番号）に対応したError Messageを返す関数
func ErrorText(code int) string {
	return errorText[code]
}

// AppError カスタムエラー用の構造体
type AppError struct {
	Code int
	Err  error
}

// NewAppError Error Codeとerrを元にAppError型のオブジェクトを生成する関数
func NewAppError(code int, err error) *AppError {
	if tae, ok := errors.Cause(err).(*AppError); ok {
		code = tae.Code
	}
	return &AppError{Code: code, Err: err}
}

// Error カスタムエラー用のメソッド
func (e *AppError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

// Wrap 任意のエラーメッセージをラップして新規エラーを返すメソッド
func (e *AppError) Wrap() (err error) {
	return NewAppError(e.Code, fmt.Errorf(ErrorText(e.Code)))
}

// --------
// main.go↓
// --------

func main() {
	e := echo.New()

	u := new(User)
	InitRouting(e, u)

	e.Start(":9020")
}

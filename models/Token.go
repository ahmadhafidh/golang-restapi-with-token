package models

type Token struct {
	AccessToken     string `json:"access_token"`
	ExpiredDatetime string `json:"expired_datetime"`
	TokenType       string `json:"token_type"`
}

type TokenBNI struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type ResponseToken struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Data            Token  `json:"data"`
}

func ResponseTokenBNI() *TokenBNI {
	return &(TokenBNI{})
}

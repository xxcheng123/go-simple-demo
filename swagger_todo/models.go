package main

type ParamUserCommon struct {
	Username string `json:"username" binding:"gte=3,lte=8,required"`
	Password string `json:"password" binding:"gte=8,lte=16,required"`
}

type UserInfo struct {
	Username string `json:"username"`
}

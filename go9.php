/**
jwt 创建 概念 解析
解析 token.Claims.()
断言成指针 *&MyClamims

 */
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 如何依赖一个结构体实现接口
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
func main(){
	mySigningKey := []byte("wosihqimiaofangguowoba")
	// MapClaims map
	// StandardClaims struct
	c := MyClaims{
		Username: "qimiao",
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60, //当前时间生效
			ExpiresAt: time.Now().Unix() + 7, //失效时间生效
			Issuer: "qimiao", //签发这
		},
	}
	t :=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	s,e :=t.SignedString(mySigningKey)
	if e!=nil{
		fmt.Println("%s",e)
	}
	fmt.Println(s)
	time.Sleep(6*time.Second) // 解析之前睡6秒
	// 收到发过来的数据解密 1.传递token 2.解析模板 3.按照规矩写func
	token,err := jwt.ParseWithClaims(s,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return  mySigningKey, nil
	})
	if err != nil{
		fmt.Printf("%s",err)
		return
	}
	fmt.Println(token)
	fmt.Println(token.Claims.(*MyClaims))  //{qimiao { 1621647400  0 qimiao 1621647333 }}

	// 头 体 加密串
	//{ 0xc00009a090 map[alg:HS256 typ:JWT] {qimiao { 1621589095  0 qimiao 1621596235 }}  false}
}


// 创建一个token
//func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString(j.SigningKey)
//}

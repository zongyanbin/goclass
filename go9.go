/**
jwt
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
			ExpiresAt: time.Now().Unix() - 60*60*2, //失效时间生效
			Issuer: "qimiao", //签发这
		},
	}
	t :=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	s,e :=t.SignedString(mySigningKey)
	if e!=nil{
		fmt.Println("%s",e)
	}
	fmt.Println(s)

	jwt.ParseWithClaims(s,MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		
	})

	// 头 体 加密串
	//{ 0xc00009a090 map[alg:HS256 typ:JWT] {qimiao { 1621589095  0 qimiao 1621596235 }}  false}
}
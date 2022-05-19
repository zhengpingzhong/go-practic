package tool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCokCmg/+yi5qzeJJ4wMdbbr42UGYJM+jB8lelOgsslTXhmBDimO0cy4JH3/5ABPy3w3ARWz4SOyqWd9Vsij4L8iGY5SsO6s7bINEktcrjEvWauBPHvqvaphHeqJQid5j8izh7jKvpdzJWrM/Xx4Kd4PRyCkvglNvGhphCEkfslV3X3uSPxfqBp0KhDYSpd7hW3uepuFvWpj+r1sSd4lmFgjDIHkuMYdOEypZK4i3M+7FzvFWiT6XJoijKeRWU2NPsLrt+JCBKUytkJvEIpX0gUpNvfyIBs7RrFfXQpK+TZcghq9pAOnB3OI1WxLR6k3PmecdqdaJsbYWj5Knq8QQ9vAgMBAAECggEAbFGHjsAmzl+xgmbBiHLegPmE30CtLM4xTj3/WMZPyL7JubTIunmpPRZ2dokxfrucyAR6fVaI9kFev+ylDcZvA9ya5AIZWOaXdHu08pGgDj0vuG/REJQqQbCwPRPjP6Rvuev+iTha/t+y5ftT9oTI2UlcAsRy1WdDSv8t6iLs4ocNVw87+Kvs/g/ysuoGiwGLDqNeiLULd/t6QAnAKztZxI3cguxnUPGjEmQoiXrkufWhPsGRGC7tHLE9nd47zUwXmmKfCFUqXvWaX8KWi3Bc4w3HPbplUDbjCTY6gdiMTiiFtWhjpp8o3FgcaSalUWnwJTOsxAlWwWP/rp3vyKqnEQKBgQDyIBlMNr+qhsTPLuNH58QrEdsNR01cTr2B3qHpoP6NNuz2J8f2RfcX/1iHzBpjmcDowfVTFIuSF4VAI/X30p0gkFBD+rGRGWWjDqgCpwxpaZUaW7k6tI9mvfF4RaFrUJqKXrXHSpEWjrmVUwCL2o+43qwojMEtAxiWXMliTqs7BQKBgQCyOO3paDKB2AagtnYeiZ18NVOB+2yhNuQMsjm01WU/YUTffzVG+VEIv1LDYQkKXRkZ0c46hMCZZ3v0cMebINlpypYZEGKfXUTv0ueqLbn5sqTOB2Yq6+BEVpAO84ijVv23XuxeUii4UwqgcmK1Mg7mYCq8jwfF89gLHXUg5e3y4wKBgF7GS8C2aX3hdOY4T0TVJsdJHPpsn45zT+Cm+a/OkfVKu6R62929c0QUcl3teaLlC0rQqSuVDjDoPgNmUCBqc5DMjG6fkaPfUhZqZT9sjgIxDukbYDzncn5f+UX/03iDO4rQGc1hlbtDsMcPj1Yu/jT6DF9vASqSXKx2s2lkq8fVAoGAU5K2ndkqthCjKwHk6fUGUJSBEFbiXcbrsbhxrwbSb6IhNdzlEt8ezMdK/CuonZ+CgcUMvAw8dLM+QSYRtAD3ctB5Ck/Qr4KBLSTtVbQTna/T8Hkvw7jQdbtR2nEBBfdpCnAMzz4fsi/er9ZzyxwqOIWoBLvdnUn3aV7q6f5oUpkCgYEAmj+4VmQ9v+52JpKJTDuT4ayBfF8te4qDiz2tQnnNguq+FY6fGyUh663ciSBugDes3w3RSR/Yin7Ndkjm58Rk2vRaLbE/udcJ/m+nz1yAoDP9r+0plKWUl9drq19mMu7UB1fXKzZ1DgmCkJ6xJeyANYslA03Zb82jjNOFfcHELnQ=
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqJApoP/souas3iSeMDHW26+NlBmCTPowfJXpToLLJU14ZgQ4pjtHMuCR9/+QAT8t8NwEVs+EjsqlnfVbIo+C/IhmOUrDurO2yDRJLXK4xL1mrgTx76r2qYR3qiUIneY/Is4e4yr6XcyVqzP18eCneD0cgpL4JTbxoaYQhJH7JVd197kj8X6gadCoQ2EqXe4Vt7nqbhb1qY/q9bEneJZhYIwyB5LjGHThMqWSuItzPuxc7xVok+lyaIoynkVlNjT7C67fiQgSlMrZCbxCKV9IFKTb38iAbO0axX10KSvk2XIIavaQDpwdziNVsS0epNz5nnHanWibG2Fo+Sp6vEEPbwIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	//priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	//解析PKCS8格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), ciphertext)
}

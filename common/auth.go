package common

import(
    "crypto/rsa"
    "io/ioutil"
    "log"
    "time"
    jwt "github.com/dgrijalva/jwt-go"
)

type AppClaims struct{
    Username string `json:"username"`
    Role string `json:"role"`
    jwt.StandardClaims
}

const(
    privKeyPath = "./keys/app.rsa"
    pubKeyPath = "./keys/app.rsa.pub"
)

var(
    verifyKey *rsa.PublicKey
    signKey *rsa.PrivateKey
)

func initKeys(){
    var err error
    signBytes, err := ioutil.ReadFile(privKeyPath)
    if err != nil{
        log.Fatalf("[initKeys]: %s\n", err)
        panic(err)
    }
    signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
    verifyBytes, err := ioutil.ReadFile(pubKeyPath)
    if err != nil{
        log.Fatalf("[initKeys]: %s\n", err)
        panic(err)
    }
    verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

func GenerateJWT(name, role string)(string, error){
    claims := AppClaims{
        name,
        role,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer: "admin",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    log.Println("token : ",token)
    ss, err := token.SignedString(signKey)
    if err != nil{
        return "", err
    }
    return ss,nil
}

package interceptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func EncryptInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ler o corpo da solicitação original
		body, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		// Criptografar o corpo usando um algoritmo de criptografia
		encryptedBody, err := encryptBody(body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		// Definir o corpo criptografado na solicitação
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(encryptedBody))
		c.Request.ContentLength = int64(len(encryptedBody))
		c.Header("Content-Type", "application/json")

		c.Next()
	}
}

// Rota de exemplo que exibe o corpo da solicitação.
func readBody(body io.Reader) string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)
	return buf.String()
}

// Função auxiliar para criptografar o corpo.
func encryptBody(body []byte) ([]byte, error) {

	// Chave de criptografia
	key := []byte("jobsstudy1234567")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(body))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], body)

	return []byte(base64.URLEncoding.EncodeToString(ciphertext)), nil
}

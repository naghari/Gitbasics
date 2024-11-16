package upload

import (
    "io"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/hex"
    "crypto/rand"
)

func encrypt(skey string, plaintext []byte) string {
    /*decode Hex key to bytes*/
    key, err := hex.DecodeString(skey)
    if err != nil {
        	panic(err)
    }
    /* Create AES cipher with a given key*/
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    return base64.StdEncoding.EncodeToString(ciphertext)
}

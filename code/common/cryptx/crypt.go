/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/2/16 9:14
 * @version     v1.0
 * @filename    crypt.go
 * @description
 ***************************************************************************/
package cryptx

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

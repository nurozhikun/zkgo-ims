package sutils

import (
	"bytes"
	"crypto/sha256"
	"time"

	"gitee.com/sienectagv/gozk/zutils"
)

func sumOfPass(buff *bytes.Buffer) []byte {
	s := sha256.Sum256(buff.Bytes())
	return s[:]
}

func SaltPassCreate(user, pass string) (sum []byte, salt string, err error) {
	buff := bytes.NewBufferString("wuzk@sie:" + time.Now().UTC().Format("20060102-15:04:05.000"))
	salt = buff.String()
	buff.WriteString(pass)
	buff.WriteString(user)
	// s := sha256.Sum256(buff.Bytes())
	// sum = s[:]
	sum = sumOfPass(buff)
	return
}

func SaltPassCheck(user, pass, salt string, sum []byte) (err error) {
	buff := bytes.NewBufferString(salt)
	buff.WriteString(pass)
	buff.WriteString(user)
	if bytes.Compare(sum, sumOfPass(buff)) == 0 {
		return nil
	} else {
		return zutils.ErrUserOrPassMiss
	}
}

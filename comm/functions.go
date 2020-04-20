package comm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"gin-sam/conf"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//当前时间戳
func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).UnixNano())
}

// 时间戳转化为dateY-m-d H:i:s
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeForm)
	} else {
		return time.Now().Format(conf.SysTimeForm)
	}
}

func FormatFromUnixTimeShort(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeFormShort)
	} else {
		return time.Now().Format(conf.SysTimeFormShort)
	}
}

// 将字符串转化为时间
func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(conf.SysTimeForm, str, conf.SysTimeLocation)
}

//获取随机数
func Random(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}
func encrypt(key, text []byte) ([]byte, error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	b := base64.StdEncoding.EncodeToString(text)
	cipherText := make([]byte, aes.BlockSize+len(b))
	iv := cipherText[:aes.BlockSize]
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))
	return cipherText, nil
}
func decrypt(key, text []byte) ([]byte, error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("cipherText too short")

	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	bytes, e := base64.StdEncoding.DecodeString(string(text))
	if e != nil {
		return nil, e
	}
	return bytes, nil
}

func AddSlashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

func StripSlashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLen := len(strRune)
	for i := 0; i < strLen; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}
func Ip4ToInt(ip string) int64 {
	bits := strings.Split(ip, ".")
	if len(bits) == 4 {
		b0, _ := strconv.Atoi(bits[0])
		b1, _ := strconv.Atoi(bits[0])
		b2, _ := strconv.Atoi(bits[0])
		b3, _ := strconv.Atoi(bits[0])
		var sum int64
		sum += int64(b0) << 24
		sum += int64(b1) << 16
		sum += int64(b2) << 8
		sum += int64(b3)
		return sum
	} else {
		return 0
	}
}

// 得到当前时间下一天的零
func NextDayDuration() time.Duration {
	year, month, day := time.Now().Add(time.Hour * 24).Date()
	next := time.Date(year, month, day, 0, 0, 0, 0, conf.SysTimeLocation)
	return next.Sub(time.Now())
}

//从接口类型安全获取到int64
func GetInt64(i interface{}, d int64) int64 {
	if i == nil {
		return d
	}
	switch i.(type) {
	case string:
		num, err := strconv.Atoi(i.(string))
		if err != nil {
			return d
		} else {
			return int64(num)
		}

	case []byte:
		bits := i.([]byte)
		if len(bits) == 8 {
			return int64(binary.LittleEndian.Uint64(bits))
		} else if len(bits) <= 4 {
			num, e := strconv.Atoi(string(bits))
			if e != nil {
				return d
			} else {
				return int64(num)
			}
		}
	case uint:
		return int64(i.(uint))
	case uint8:
		return int64(i.(uint8))
	case uint16:
		return int64(i.(uint16))
	case uint32:
		return int64(i.(uint32))
	case uint64:
		return int64(i.(uint64))
	case int:
		return int64(i.(int))
	case int8:
		return int64(i.(int8))
	case int32:
		return int64(i.(int32))
	case int64:
		return i.(int64)
	case float32:
		return int64(i.(float32))
	case float64:
		return int64(i.(float64))
	}
	return d
}
func GetString(str interface{}, d string) string {
	if str == nil {
		return d
	}
	switch str.(type) {
	case string:
		return str.(string)
	case []byte:
		return string(str.([]byte))
	}
	return fmt.Sprintf("%s", str)
}

func GetInt64FromMap(dm map[string]interface{}, key string, dft int64) int64 {
	i, ok := dm[key]
	if !ok {
		return dft
	}
	return GetInt64(i, dft)
}

func GetInt64FromStringMap(dm map[string]string, key string, dft int64) int64 {
	data, ok := dm[key]
	if !ok {
		return dft
	}
	return GetInt64(data, dft)
}

func GetStringFromMap(dm map[string]interface{}, key, dft string) string {
	data, ok := dm[key]
	if !ok {
		return dft
	}
	return GetString(data, dft)
}
func GetStringFromStringMap(dm map[string]string, key string, dft string) string {
	data, ok := dm[key]
	if !ok {
		return dft
	}
	return data
}

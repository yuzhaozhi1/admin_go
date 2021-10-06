package utils

import (
	"errors"
	"reflect"
	"strings"
)

// validator 验证器

type Rules map[string][]string // 路由

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

// RegisterRule 注册自定义规则方案, 建议在路由初始化层即注册
func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap != nil {
		return errors.New(key + "已注册, 无法重复注册!")
	} else {
		CustomizeMap[key] = rule
		return
	}
}

// NotEmpty 非空 不能为其对应类型的0值
func NotEmpty() string {
	return "notEmpty"
}

func Verify(st interface{}, realMap Rules) (err error) {
	// realMap LoginVerify = Rules{
	// "CaptchaId": {NotEmpty()}, "Captcha":{NotEmpty()}, "Username":{NotEmpty()}, "Password":{NotEmpty()}}

	// 比较map
	compareMpa := map[string]bool{
		"lt": true, // 小于
		"le": true, // 小于等于
		"eq": true, // 等于
		"ne": true, // 不等于
		"ge": true, // 大于等于
		"gt": true, // 大于
	}
	typ := reflect.TypeOf(st)  // 获取 类型
	val := reflect.ValueOf(st) // 获取值

	kd := typ.Kind() // 返回原始的类型
	if kd != reflect.Struct {
		return errors.New("expect struct, 不是结构体")
	}

	num := val.NumField() // 该结构体一共有多少个字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i) // 获取单条的字段
		val := val.Field(i)    // 获取单条字段的值
		if len(realMap[tagVal.Name]) > 0 {
			// Rules{"CaptchaId": {NotEmpty()}, "Captcha":{NotEmpty()}, "Username":{NotEmpty()}, "Password":{NotEmpty()}}
			for _, ToBecheckField := range realMap[tagVal.Name] {
				switch {
				case ToBecheckField == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + ":该字段的值不可为空")
					}
				case compareMpa[strings.Split(ToBecheckField, "=")[0]]:
					if !compareVerify(val, ToBecheckField) {

					}

				}
			}
		}

	}
}

// isBlank 非空校验
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	// 否则为map ,数组, 切片类型, 与其为0 值的数据进行比较
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// compareVerify 长度和数字的校验方法, 根据类型自动校验
func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	
	}

}

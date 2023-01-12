package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key string
	Msg string
}

func (v *ValidError) Error() string {
	return v.Msg
}

type ValidErrors []*ValidError

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	s := make([]string, 0)
	for _, e := range v {
		s = append(s, e.Error())
	}
	return s
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var validErr ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		trans, _ := c.Value("trans").(ut.Translator)
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, validErr
		}
		for k, v := range err.Translate(trans) {
			validErr = append(validErr, &ValidError{
				Key: k,
				Msg: v,
			})
		}
		return false, validErr
	}
	return true, nil
}

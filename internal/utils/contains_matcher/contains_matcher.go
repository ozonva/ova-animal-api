package contains_matcher

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"reflect"
)

func New(x interface{}) gomock.Matcher {
	matcher := containsMatcher{x}
	return &matcher
}

type containsMatcher struct {
	element interface{}
}

func (this *containsMatcher) Matches(src interface{}) bool {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(src)

		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == this.element {
				return true
			}
		}
	default:
		panic("works with slices only")
	}

	return false
}

func (this *containsMatcher) String() string {
	return fmt.Sprintf("Contains %s", this.element)
}

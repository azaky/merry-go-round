package b

import (
	"github.com/azaky/merry-go-round/global"
)

type Service struct {
	ServiceB global.ServiceB
}

func (s *Service) GetName() string {
	return "Merry"
}

func (s *Service) GetSomething() string {
	return "Something: " + s.ServiceB.GetSomething()
}

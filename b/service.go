package b

import (
	"github.com/azaky/merry-go-round/global"
)

type Service struct {
	ServiceA global.ServiceA
}

func (s *Service) GetName() string {
	return "name: " + s.ServiceA.GetName()
}

func (s *Service) GetSomething() string {
	return "else"
}

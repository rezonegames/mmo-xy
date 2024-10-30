package util

import (
	"mmo-xy/models"
	"mmo-xy/pkg/z"

	"github.com/lonng/nano/session"
)

// BindUser todo：如果是集群模式，使用remote方法同步player数据，并保存到session
func BindUser(s *session.Session, p *models.User) error {
	err := s.Bind(p.UserId)
	if err != nil {
		return err
	}
	s.Set("user", p)
	return nil
}

func GetPlayer(s *session.Session) (*models.User, error) {
	var (
		user = s.Value("user")
	)
	if v, ok := user.(*models.User); ok {
		return v, nil
	}
	return nil, z.NilError{}
}

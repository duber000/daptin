package resource

import (
  log "github.com/Sirupsen/logrus"
  "github.com/artpar/gocms/server/auth"
)

type Permission struct {
  UserId      string `json:"user_id"`
  UserGroupId []auth.GroupPermission `json:"usergroup_id"`
  Permission  int64 `json:"permission"`
}

func (p Permission) CanExecute(userId string, usergroupId []auth.GroupPermission) bool {
  return p.CheckBit(userId, usergroupId, 1)
}

func (p Permission) CanWrite(userId string, usergroupId []auth.GroupPermission) bool {
  return p.CheckBit(userId, usergroupId, 2)
}

func (p Permission) CanRead(userId string, usergroupId []auth.GroupPermission) bool {
  return p.CheckBit(userId, usergroupId, 4)
}


func (p1 Permission) CheckBit(userId string, usergroupId []auth.GroupPermission, bit int64) bool {
  if userId == p1.UserId {
    p := p1.Permission / 100
    log.Infof("Check against user: %v", p)
    return (p & bit) == bit
  }

  for _, uid := range usergroupId {

    for _, gid := range p1.UserGroupId {
      if uid.ReferenceId == gid.ReferenceId {
        p := (gid.Permission % 100) / 10
        p = p % 10
        log.Infof("Check against group [%v]: %v", gid.ReferenceId, p)
        return (p & bit) == bit
      }
    }
  }

  p := p1.Permission % 10
  log.Infof("Check against world: %v == %v", p, (p & bit) == bit)
  return (p & bit) == bit
}

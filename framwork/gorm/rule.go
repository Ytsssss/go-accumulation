package gorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

// Rule 实体
type Rule struct {
	Id         int64     `gorm:"id pk autoincr"`
	Response   string    `gorm:"response"`
	Hit        int64     `gorm:"hit"`
	Name       string    `gorm:"name"`
	UriId      int64     `gorm:"uri_id uniqueIndex:uri_id_identity_unique"`
	RuleStatus int       `gorm:"rule_status"`
	Identity   string    `gorm:"identity uniqueIndex:uri_id_identity_unique"`
	Parameters string    `gorm:"parameters"`
	CreateTime time.Time `gorm:"<-"`
	UpdateTime time.Time `gorm:"<-"`
}

var (
	ruleDAOInitOnce sync.Once
	ruleDAOInstance RuleDAO
	sqliteDB        *gorm.DB
)

type RuleDAOImpl struct {
	*gorm.DB
}

type RuleDAO interface {
	InsertWithDuplicate(rules []*Rule) (int64, error)
}

func (r *RuleDAOImpl) InsertWithDuplicate(rules []*Rule) (int64, error) {
	create := r.DB.Debug().Table("t_rule").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uri_id"}, {Name: "identity"}},
		DoUpdates: clause.AssignmentColumns([]string{"response", "hit", "name", "rule_status", "parameters"}),
	}).Create(rules)

	return create.RowsAffected, create.Error
}

func Init() error {
	db, err := gorm.Open(sqlite.Open("123.db3"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	sqliteDB = db
	return err
}

// GetRuleDaoIns 获取dao单例
var GetRuleDaoIns = func() RuleDAO {
	ruleDAOInitOnce.Do(func() {
		ruleDAOInstance = &RuleDAOImpl{sqliteDB}
	})

	return ruleDAOInstance
}

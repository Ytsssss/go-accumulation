package dbsync

import (
	"fmt"
	"log"

	"github.com/Ytsssss/go_accumulation/utils"
)

var (
	dumpSchema = "mysqldump  --set-gtid-purged=OF --column-statistics=0 --no-data -h %s -P %d -u %s -p%s %s %s > %s"
	dumpData   = "mysqldump  -t --set-gtid-purged=OF --column-statistics=0 -h %s -P %d -u %s -p%s %s %s --where=\"%s\" > %s"
	importData = "mysql -h %s -P %d -u %s -p%s %s < %s"
)

type MysqlConn struct {
	IP         string
	Port       int
	User       string
	Password   string
	Database   string
	Table      string
	Where      string
	SchemaPath string
	DataPath   string
}

func (conn *MysqlConn) DumpSchema() error {
	dumpBash := fmt.Sprintf(dumpSchema, conn.IP, conn.Port, conn.User, conn.Password, conn.Database, conn.Table, conn.SchemaPath)
	command, err := utils.RunBashCommand(dumpBash)
	if err != nil {
		return err
	}
	log.Printf(command)
	return nil
}

func (conn *MysqlConn) DumpData() error {
	dumpBash := fmt.Sprintf(dumpData, conn.IP, conn.Port, conn.User, conn.Password, conn.Database, conn.Table, conn.Where, conn.DataPath)
	command, err := utils.RunBashCommand(dumpBash)
	if err != nil {
		return err
	}
	log.Printf(command)
	return nil
}

func (conn *MysqlConn) Import() error {
	dumpBash := fmt.Sprintf(importData, conn.IP, conn.Port, conn.User, conn.Password, conn.Database, conn.DataPath)
	command, err := utils.RunBashCommand(dumpBash)
	if err != nil {
		return err
	}
	log.Printf(command)
	return nil
}

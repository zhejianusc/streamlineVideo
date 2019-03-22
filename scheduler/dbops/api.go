package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// api -> videoid - mysql
// dispatcher -> mysql-videoid -> datachannel
// executor -> datachannel-videoid -> delete videos
func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
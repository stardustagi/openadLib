package mysql

import (
	"testing"
	"xorm.io/builder"
)

type TestTrans struct {
	Id int `json:"id" xorm:"id"`
}

type TestTrans1 struct {
	Id    int `json:"id" xorm:"id"`
	Value int `json:"value" xorm:"value"`
}

func (m *TestTrans) TableName() string {
	return "test_trans"
}
func TestEntityNew_Create(t *testing.T) {

	//logger := hclog.Default()
	c := Config{
		ShowSql: true,
		Master:  "root:123456@tcp(127.0.0.1)/ad-aws?charset=utf8mb4&parseTime=true&loc=Local",
	}
	conn, _ := NewConn(&c, "test.log")
	createTable := `create table if not exists test_trans(id int)engine=innodb`
	conn.Query(createTable)
	dao := NewBaseDao(conn)
	session := dao.NewSession()
	defer session.Close()
	session.Begin()
	t1 := TestTrans{Id: 1}
	t2 := TestTrans{Id: 2}
	t3 := TestTrans{Id: 3}
	t4 := TestTrans1{Id: 3, Value: 1}
	_, err := session.InsertOne(&t1)

	if err != nil {
		t.Fatal(err)
		session.Rollback()
		return
	}
	session.Begin()
	session.InsertOne(&t2)
	if _, err := dao.InsertOne(&t3); err != nil {
		t.Fatal(err)
		return
	}

	s, m, e := builder.ToSQL(builder.Expr("value + ?", 3))
	if e != nil {
		t.Fatal(e)
		return
	}
	t.Log(s, m)
	_, err = session.Where("id = ?", 1).Update(&t4, builder.Expr("value + ?", 3))
	if err != nil {
		t.Fatal(err)
		return
	}

	session.Commit()
}

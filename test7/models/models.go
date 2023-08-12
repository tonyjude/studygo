package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Store struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Customer struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(Store), new(Customer))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")

	dbcon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	fmt.Print(dbcon)
	orm.RegisterDataBase("default", "mysql", dbcon /*"root:root@tcp(localhost:3306)/test?charset=utf8"*/) //密码为空格式
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	created := time.Now()
	topicTime := time.Now()
	cate := &Store{Title: name, Created: created, TopicTime: topicTime}

	qs := o.QueryTable("store")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllCategories() ([]*Store, error) {
	o := orm.NewOrm()

	cates := make([]*Store, 0)

	qs := o.QueryTable("store")
	_, err := qs.All(&cates)

	return cates, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Store{Id: cid}
	_, err = o.Delete(cate)

	return err
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Customer{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}

func GetAllTopics(isDesc bool) ([]*Customer, error) {
	o := orm.NewOrm()

	topics := make([]*Customer, 0)

	qs := o.QueryTable("customer")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-id").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(tid string) (*Customer, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Customer)

	qs := o.QueryTable("customer")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	return topic, err
}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Customer{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}

	return nil
}

func DelTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Customer{Id: tid}
	_, err = o.Delete(topic)

	return err
}

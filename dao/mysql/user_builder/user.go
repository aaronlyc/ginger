package user_builder

import (
	"context"
	"errors"
	"github.com/gofuncchan/ginger/dao/mysql"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"time"
)

/*
   This code is generated by gendry
   gforge dao -u fileservice_user -p 123456 -d fileservice_db -t user
*/

// User is a mapping object for user table in mysql
type User struct {
	Age      byte      `ddb:"age"`
	Avatar   string    `ddb:"avatar"`
	CreateAt time.Time `ddb:"create_at"`
	Email    string    `ddb:"email"`
	Gender   int8      `ddb:"gender"`
	ID       uint      `ddb:"id"`
	Name     string    `ddb:"name"`
	Password string    `ddb:"password"`
	Phone    string    `ddb:"phone"`
	Salt     string    `ddb:"salt"`
	Status   int8      `ddb:"status"`
	UpdateAt time.Time `ddb:"update_at"`
}

const (
	TableName = "user"
)

// GetOne gets one record from table user by condition "where"
func GetOne(where map[string]interface{}) (*User, error){
	if nil == mysql.Db {
		return nil, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect(TableName, where, nil)
	if nil != err {
		return nil, err
	}

	row, err := mysql.Db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()
	var res *User
	err = scanner.Scan(row, &res)
	return res, err
}

// GetMulti gets multiple records from table user by condition "where"
func GetMulti(where map[string]interface{}) ([]*User, error) {
	if nil == mysql.Db {
		return nil, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect(TableName, where, nil)
	if nil != err {
		return nil, err
	}
	row, err := mysql.Db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()
	var res []*User
	err = scanner.Scan(row, &res)
	return res, err
}

// Insert inserts an array of data into table user
func Insert(data []map[string]interface{}) (int64, error) {
	if nil == mysql.Db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildInsert(TableName, data)
	if nil != err {
		return 0, err
	}
	result, err := mysql.Db.Exec(cond, vals...)
	if nil != err || nil == result {
		return 0, err
	}
	return result.LastInsertId()
}

// Update updates the table user
func Update(where, data map[string]interface{}) (int64, error) {
	if nil == mysql.Db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildUpdate(TableName, where, data)
	if nil != err {
		return 0, err
	}
	result, err := mysql.Db.Exec(cond, vals...)
	if nil != err {
		return 0, err
	}
	return result.RowsAffected()
}

// Delete deletes matched records in user
func Delete(where, data map[string]interface{}) (int64, error) {
	if nil == mysql.Db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildDelete(TableName, where)
	if nil != err {
		return 0, err
	}
	result, err := mysql.Db.Exec(cond, vals...)
	if nil != err {
		return 0, err
	}
	return result.RowsAffected()
}

// GetCount
func GetCount(where map[string]interface{}) (int64, error) {
	if nil == mysql.Db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	res, err := builder.AggregateQuery(context.TODO(), mysql.Db, TableName, where, builder.AggregateCount("*"))
	if nil != err {
		return 0, err
	}

	return res.Int64(), err
}
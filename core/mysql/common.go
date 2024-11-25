package mysql

import (
	"fmt"
	"github.com/stardustagi/openadLib/core/errors"
	"github.com/stardustagi/openadLib/core/logger"
)

func DoGet(call func() (bool, error)) error {
	has, err := call()
	if nil != err {
		return err
	}
	if !has {
		return ErrGetEmpty
	}
	return nil
}
func DoUpdate(call func() (int64, error)) error {
	ret, err := call()
	if nil != err {
		return err
	}
	if 0 == ret {
		return ErrUpdatedEmpty
	}
	return nil
}

func DoInsert(call func() (int64, error)) error {
	ret, err := call()
	if nil != err {
		return err
	}
	if 0 == ret {
		return ErrInsertedEmpty
	}
	return nil
}

func DoDelete(call func() (int64, error)) error {
	ret, err := call()
	if nil != err {
		return err
	}
	if 0 == ret {
		return ErrDeletedEmpty
	}
	return nil
}

type SessionWrapper struct {
	dao     BaseDao
	session SessionDao
	err     errors.StackError
}

type ExecuteFunc func(session SessionDao) errors.StackError

func NewSessionWrapper(s BaseDao) SessionWrapper {
	sw := SessionWrapper{
		dao:     s,
		session: s.NewSession(),
	}
	if err := sw.session.Begin(); err != nil {
		sw.err.e
	}
	return sw
}

func (sw SessionWrapper) Execute(call ExecuteFunc) SessionWrapper {
	defer func() {
		if r := recover(); r != nil {
			// 这里可以记录日志或执行其他操作
			fmt.Println("Recovered from panic in SessionWrapper Execute:", r)
			if err := sw.session.Rollback(); err != nil {
				logger.Error("Rollback error: ", err)
			}
		}
	}()
	// 有错误直接返回
	if sw.err != nil {
		return sw
	}
	if sw.err = call(sw.session); sw.err != nil {
		logger.Error("Execute error: ", sw.err)
		_ = sw.session.Rollback()
	}
	return sw
}

func (sw SessionWrapper) CommitAndClose() error {
	defer sw.session.Close()
	if sw.err == nil {
		sw.err = sw.Commit()
	}
	return sw.err
}

func (sw SessionWrapper) Commit() error {
	// 没有错误，提交commit
	if sw.err == nil {
		sw.err = sw.session.Commit()
	} else {
		// 有错误，回滚
		_ = sw.session.Rollback()
	}

	return sw.err
}

func (sw SessionWrapper) Close() error {
	// 有错误则回滚后更新
	if sw.err != nil {
		_ = sw.session.Rollback()
	}
	sw.session.Close()
	return sw.err
}

func (sw SessionWrapper) GetErr() error {
	return sw.err
}

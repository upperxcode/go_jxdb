package jxdb

import (
	"fmt"
	"strings"

	"github.com/upperxcode/go_jxdb/pkg/sql"
)

type DbControl interface {
	BeforeInsert(model interface{}) error
	AfterInsert(model interface{}) error
	BeforeUpdate(model interface{}) error
	AfterUpdate(model interface{}) error
	BeforeDelete(id int) error
	AfterDelete(id int) error
}

type GenRepository[T any] struct {
	Db         sql.Database
	TableName  string
	Fields     []string
	ScanFunc   func(interface{}) (T, error)
	ValuesFunc func(T) []interface{}
	IDValue    func(T) interface{}
	Control    DbControl
	Order      string
	Limit      int
	Joins      []string
}

func (r *GenRepository[T]) FindByID(id int) (T, error) {
	var result T
	query := fmt.Sprintf("SELECT %s FROM %s %s WHERE %s.id = $1", r.getFields(), r.TableName, r.getJoins(), r.TableName)
	println("ID => ", id)
	err := r.Db.Get(&result, query, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *GenRepository[T]) FindAll() ([]T, error) {
	var results []T
	query := fmt.Sprintf("SELECT %s FROM %s %s %s %s", r.getFields(), r.TableName, r.getJoins(), r.getOrder(), r.getLimit())
	println(query)
	err := r.Db.Select(&results, query)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *GenRepository[T]) Find(where string, args ...interface{}) ([]T, error) {
	var results []T
	query := fmt.Sprintf("SELECT %s FROM %s %s WHERE %s %s %s", r.getFields(), r.TableName, r.getJoins(), where, r.getOrder(), r.getLimit())
	err := r.Db.Select(&results, query, args...)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *GenRepository[T]) Insert(model T) error {
	if r.Control != nil {
		if err := r.Control.BeforeInsert(model); err != nil {
			return err
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", r.TableName, r.getFields(), r.getPlaceholders())
	err := r.Db.Insert(query, r.ValuesFunc(model)...)
	if err != nil {
		return err
	}

	if r.Control != nil {
		if err := r.Control.AfterInsert(model); err != nil {
			return err
		}
	}
	return nil
}

func (r *GenRepository[T]) Update(model T) error {
	if r.Control != nil {
		if err := r.Control.BeforeUpdate(model); err != nil {
			return err
		}
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", r.TableName, r.getUpdateFields())
	err := r.Db.Update(query, append(r.ValuesFunc(model), r.IDValue(model))...)
	if err != nil {
		return err
	}

	if r.Control != nil {
		if err := r.Control.AfterUpdate(model); err != nil {
			return err
		}
	}
	return nil
}

func (r *GenRepository[T]) Delete(id int) error {
	if r.Control != nil {
		if err := r.Control.BeforeDelete(id); err != nil {
			return err
		}
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.TableName)
	err := r.Db.Delete(query, id)
	if err != nil {
		return err
	}

	if r.Control != nil {
		if err := r.Control.AfterDelete(id); err != nil {
			return err
		}
	}
	return nil
}

func (r *GenRepository[T]) getFields() string {
	return strings.Join(r.Fields, ", ")
}

func (r *GenRepository[T]) getPlaceholders() string {
	placeholders := make([]string, len(r.Fields))
	for i := range r.Fields {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ", ")
}

func (r *GenRepository[T]) getUpdateFields() string {
	updateFields := make([]string, len(r.Fields))
	for i, field := range r.Fields {
		updateFields[i] = fmt.Sprintf("%s = ?", field)
	}
	return strings.Join(updateFields, ", ")
}

func (r *GenRepository[T]) getOrder() string {
	if r.Order != "" {
		return r.Order
	}
	return ""
}

func (r *GenRepository[T]) getLimit() string {
	if r.Limit > 0 {
		return fmt.Sprintf("LIMIT %d", r.Limit)
	}
	return ""
}

func (r *GenRepository[T]) getJoins() string {
	if len(r.Joins) > 0 {
		return strings.Join(r.Joins, " ")
	}
	return ""
}

package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
	sq "gopkg.in/Masterminds/squirrel.v1"
	"gopkg.in/gorp.v1"
)

type SqlStorage struct {
	executor gorp.SqlExecutor
}

func NewSQLStorage(driver string, dataSourceName string) (*SqlStorage, error) {
	db, err := sql.Open(driver, dataSourceName)

	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{
		Db: db,
	}

	if dialect, ok := migrate.MigrationDialects[driver]; ok {
		dbMap.Dialect = migrate.MigrationDialects[driver]
	} else {
		return nil, errors.New("Storage driver " + driver + " not found")
	}

	return &SqlStorage{
		executor: dbMap,
	}, nil
}

func (s *SqlStorage) CreateTablesIfNotExists() error {
	return s.executor.(*gorp.DbMap).CreateTablesIfNotExists()
}

func (s *SqlStorage) SetTypeConverter(converter gorp.TypeConverter) {
	s.executor.(*gorp.DbMap).TypeConverter = converter
}

func (s *SqlStorage) AddTableWithName(i interface{}, name string) *gorp.TableMap {
	return s.executor.(*gorp.DbMap).AddTableWithName(i, name)
}

func (s *SqlStorage) Begin() (*SqlStorage, error) {
	transaction, err := s.executor.(*gorp.DbMap).Begin()
	if err != nil {
		return nil, err
	}

	return &SqlStorage{
		executor: transaction,
	}, nil
}

func (s *SqlStorage) Commit() error {
	transaction, ok := s.executor.(*gorp.Transaction)
	if !ok {
		return errors.New("Executor is not grop.Transaction")
	}

	return transaction.Commit()
}

func (s *SqlStorage) Rollback() error {
	transaction, ok := s.executor.(*gorp.Transaction)
	if !ok {
		return errors.New("Executor is not grop.Transaction")
	}

	return transaction.Rollback()
}

func (s *SqlStorage) SelectByQuery(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	data, err := s.executor.Select(i, query, args...)
	if err != nil {
		return data, errors.New("Error getting collection from DB, query: '" + query + "', error: '" + err.Error() + "'")
	}
	return data, nil
}

func (s *SqlStorage) Select(i interface{}, builder *sq.SelectBuilder) ([]interface{}, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.SelectByQuery(i, query, args...)
}

func (s *SqlStorage) SelectOneByQuery(holder interface{}, query string, args ...interface{}) error {
	err := s.executor.SelectOne(holder, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("Error getting value from DB, query: '" + query + "', error: '" + err.Error() + "'")
	}
	return err
}

func (s *SqlStorage) SelectOne(holder interface{}, builder *sq.SelectBuilder) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.SelectOneByQuery(holder, query, args...)
}

func (s *SqlStorage) Insert(list ...interface{}) error {
	if err := s.executor.Insert(list...); err != nil {
		return errors.New("Error inserting data into DB, error: '" + err.Error() + "'")
	}

	return nil
}

func (s *SqlStorage) Update(list ...interface{}) (int64, error) {
	count, err := s.executor.Update(list...)

	if err != nil {
		err = errors.New("Error updating data in DB, error: '" + err.Error() + "'")
	}

	return count, err
}

func (s *SqlStorage) Delete(list ...interface{}) (int64, error) {
	count, err := s.executor.Delete(list...)

	if err != nil {
		err = errors.New("Error deleting data in DB, error: '" + err.Error() + "'")
	}

	return count, err
}

func (s *SqlStorage) Prepare(query string) (*sql.Stmt, error) {
	if transaction, ok := s.executor.(*gorp.Transaction); ok {
		return transaction.Prepare(query)
	} else if dbMap, ok := s.executor.(*gorp.DbMap); ok {
		return dbMap.Prepare(query)
	}

	return nil, errors.New("Executor is not grop.Transaction or gorp.DbMap")
}

func (s *SqlStorage) ExecByQuery(query string, args ...interface{}) (sql.Result, error) {
	result, err := s.executor.Exec(query, args...)
	if err != nil {
		return result, errors.New("Error executing DB query, query: '" + query + "', error: '" + err.Error() + "'")
	}
	return result, nil
}

func (s *SqlStorage) Exec(query interface{}, args ...interface{}) (sql.Result, error) {
	if b, ok := query.(*sq.SelectBuilder); ok {
		return s.ExecSelect(b)
	}

	if b, ok := query.(*sq.InsertBuilder); ok {
		return s.ExecInsert(b)
	}

	if b, ok := query.(*sq.UpdateBuilder); ok {
		return s.ExecUpdate(b)
	}

	if b, ok := query.(*sq.DeleteBuilder); ok {
		return s.ExecDelete(b)
	}

	if b, ok := query.(*sq.CaseBuilder); ok {
		return s.ExecCase(b)
	}

	return nil, errors.New("could not prepare SQL query")
}

func (s *SqlStorage) ExecSelect(builder *sq.SelectBuilder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.ExecByQuery(query, args...)
}

func (s *SqlStorage) ExecInsert(builder *sq.InsertBuilder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.ExecByQuery(query, args...)
}

func (s *SqlStorage) ExecUpdate(builder *sq.UpdateBuilder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.ExecByQuery(query, args...)
}

func (s *SqlStorage) ExecDelete(builder *sq.DeleteBuilder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.ExecByQuery(query, args...)
}

func (s *SqlStorage) ExecCase(builder *sq.CaseBuilder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("could not prepare SQL query, error: '" + err.Error() + "'")
	}
	return s.ExecByQuery(query, args...)
}

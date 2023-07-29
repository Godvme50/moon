// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"prometheus-manager/dal/model"
)

func newPromCombo(db *gorm.DB, opts ...gen.DOOption) promCombo {
	_promCombo := promCombo{}

	_promCombo.promComboDo.UseDB(db, opts...)
	_promCombo.promComboDo.UseModel(&model.PromCombo{})

	tableName := _promCombo.promComboDo.TableName()
	_promCombo.ALL = field.NewAsterisk(tableName)
	_promCombo.ID = field.NewInt32(tableName, "id")
	_promCombo.Name = field.NewString(tableName, "name")
	_promCombo.Remark = field.NewString(tableName, "remark")
	_promCombo.CreatedAt = field.NewTime(tableName, "created_at")
	_promCombo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promCombo.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_promCombo.fillFieldMap()

	return _promCombo
}

type promCombo struct {
	promComboDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String // 套餐名称
	Remark    field.String // 套餐说明
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Int64

	fieldMap map[string]field.Expr
}

func (p promCombo) Table(newTableName string) *promCombo {
	p.promComboDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promCombo) As(alias string) *promCombo {
	p.promComboDo.DO = *(p.promComboDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promCombo) updateTableName(table string) *promCombo {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Name = field.NewString(table, "name")
	p.Remark = field.NewString(table, "remark")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewInt64(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *promCombo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promCombo) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p promCombo) clone(db *gorm.DB) promCombo {
	p.promComboDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promCombo) replaceDB(db *gorm.DB) promCombo {
	p.promComboDo.ReplaceDB(db)
	return p
}

type promComboDo struct{ gen.DO }

type IPromComboDo interface {
	gen.SubQuery
	Debug() IPromComboDo
	WithContext(ctx context.Context) IPromComboDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromComboDo
	WriteDB() IPromComboDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromComboDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromComboDo
	Not(conds ...gen.Condition) IPromComboDo
	Or(conds ...gen.Condition) IPromComboDo
	Select(conds ...field.Expr) IPromComboDo
	Where(conds ...gen.Condition) IPromComboDo
	Order(conds ...field.Expr) IPromComboDo
	Distinct(cols ...field.Expr) IPromComboDo
	Omit(cols ...field.Expr) IPromComboDo
	Join(table schema.Tabler, on ...field.Expr) IPromComboDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromComboDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromComboDo
	Group(cols ...field.Expr) IPromComboDo
	Having(conds ...gen.Condition) IPromComboDo
	Limit(limit int) IPromComboDo
	Offset(offset int) IPromComboDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromComboDo
	Unscoped() IPromComboDo
	Create(values ...*model.PromCombo) error
	CreateInBatches(values []*model.PromCombo, batchSize int) error
	Save(values ...*model.PromCombo) error
	First() (*model.PromCombo, error)
	Take() (*model.PromCombo, error)
	Last() (*model.PromCombo, error)
	Find() ([]*model.PromCombo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromCombo, err error)
	FindInBatches(result *[]*model.PromCombo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromCombo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromComboDo
	Assign(attrs ...field.AssignExpr) IPromComboDo
	Joins(fields ...field.RelationField) IPromComboDo
	Preload(fields ...field.RelationField) IPromComboDo
	FirstOrInit() (*model.PromCombo, error)
	FirstOrCreate() (*model.PromCombo, error)
	FindByPage(offset int, limit int) (result []*model.PromCombo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromComboDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	WhereID(ctx context.Context, id uint) (result *model.PromCombo, err error)
}

// select * from @@table where id = @id
func (p promComboDo) WhereID(ctx context.Context, id uint) (result *model.PromCombo, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("select * from prom_combos where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p promComboDo) Debug() IPromComboDo {
	return p.withDO(p.DO.Debug())
}

func (p promComboDo) WithContext(ctx context.Context) IPromComboDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promComboDo) ReadDB() IPromComboDo {
	return p.Clauses(dbresolver.Read)
}

func (p promComboDo) WriteDB() IPromComboDo {
	return p.Clauses(dbresolver.Write)
}

func (p promComboDo) Session(config *gorm.Session) IPromComboDo {
	return p.withDO(p.DO.Session(config))
}

func (p promComboDo) Clauses(conds ...clause.Expression) IPromComboDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promComboDo) Returning(value interface{}, columns ...string) IPromComboDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promComboDo) Not(conds ...gen.Condition) IPromComboDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promComboDo) Or(conds ...gen.Condition) IPromComboDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promComboDo) Select(conds ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promComboDo) Where(conds ...gen.Condition) IPromComboDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promComboDo) Order(conds ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promComboDo) Distinct(cols ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promComboDo) Omit(cols ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promComboDo) Join(table schema.Tabler, on ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promComboDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promComboDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promComboDo) Group(cols ...field.Expr) IPromComboDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promComboDo) Having(conds ...gen.Condition) IPromComboDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promComboDo) Limit(limit int) IPromComboDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promComboDo) Offset(offset int) IPromComboDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promComboDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromComboDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promComboDo) Unscoped() IPromComboDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promComboDo) Create(values ...*model.PromCombo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promComboDo) CreateInBatches(values []*model.PromCombo, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promComboDo) Save(values ...*model.PromCombo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promComboDo) First() (*model.PromCombo, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromCombo), nil
	}
}

func (p promComboDo) Take() (*model.PromCombo, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromCombo), nil
	}
}

func (p promComboDo) Last() (*model.PromCombo, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromCombo), nil
	}
}

func (p promComboDo) Find() ([]*model.PromCombo, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromCombo), err
}

func (p promComboDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromCombo, err error) {
	buf := make([]*model.PromCombo, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promComboDo) FindInBatches(result *[]*model.PromCombo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promComboDo) Attrs(attrs ...field.AssignExpr) IPromComboDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promComboDo) Assign(attrs ...field.AssignExpr) IPromComboDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promComboDo) Joins(fields ...field.RelationField) IPromComboDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promComboDo) Preload(fields ...field.RelationField) IPromComboDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promComboDo) FirstOrInit() (*model.PromCombo, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromCombo), nil
	}
}

func (p promComboDo) FirstOrCreate() (*model.PromCombo, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromCombo), nil
	}
}

func (p promComboDo) FindByPage(offset int, limit int) (result []*model.PromCombo, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p promComboDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promComboDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promComboDo) Delete(models ...*model.PromCombo) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promComboDo) withDO(do gen.Dao) *promComboDo {
	p.DO = *do.(*gen.DO)
	return p
}

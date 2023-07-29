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

func newPromNodeDirFileGroup(db *gorm.DB, opts ...gen.DOOption) promNodeDirFileGroup {
	_promNodeDirFileGroup := promNodeDirFileGroup{}

	_promNodeDirFileGroup.promNodeDirFileGroupDo.UseDB(db, opts...)
	_promNodeDirFileGroup.promNodeDirFileGroupDo.UseModel(&model.PromNodeDirFileGroup{})

	tableName := _promNodeDirFileGroup.promNodeDirFileGroupDo.TableName()
	_promNodeDirFileGroup.ALL = field.NewAsterisk(tableName)
	_promNodeDirFileGroup.ID = field.NewInt32(tableName, "id")
	_promNodeDirFileGroup.Name = field.NewString(tableName, "name")
	_promNodeDirFileGroup.Remark = field.NewString(tableName, "remark")
	_promNodeDirFileGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_promNodeDirFileGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promNodeDirFileGroup.DeletedAt = field.NewInt64(tableName, "deleted_at")
	_promNodeDirFileGroup.FileID = field.NewInt32(tableName, "file_id")

	_promNodeDirFileGroup.fillFieldMap()

	return _promNodeDirFileGroup
}

type promNodeDirFileGroup struct {
	promNodeDirFileGroupDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String // 策略组名称
	Remark    field.String // 说明
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Int64
	FileID    field.Int32 // files id

	fieldMap map[string]field.Expr
}

func (p promNodeDirFileGroup) Table(newTableName string) *promNodeDirFileGroup {
	p.promNodeDirFileGroupDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promNodeDirFileGroup) As(alias string) *promNodeDirFileGroup {
	p.promNodeDirFileGroupDo.DO = *(p.promNodeDirFileGroupDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promNodeDirFileGroup) updateTableName(table string) *promNodeDirFileGroup {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Name = field.NewString(table, "name")
	p.Remark = field.NewString(table, "remark")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewInt64(table, "deleted_at")
	p.FileID = field.NewInt32(table, "file_id")

	p.fillFieldMap()

	return p
}

func (p *promNodeDirFileGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promNodeDirFileGroup) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["file_id"] = p.FileID
}

func (p promNodeDirFileGroup) clone(db *gorm.DB) promNodeDirFileGroup {
	p.promNodeDirFileGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promNodeDirFileGroup) replaceDB(db *gorm.DB) promNodeDirFileGroup {
	p.promNodeDirFileGroupDo.ReplaceDB(db)
	return p
}

type promNodeDirFileGroupDo struct{ gen.DO }

type IPromNodeDirFileGroupDo interface {
	gen.SubQuery
	Debug() IPromNodeDirFileGroupDo
	WithContext(ctx context.Context) IPromNodeDirFileGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromNodeDirFileGroupDo
	WriteDB() IPromNodeDirFileGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromNodeDirFileGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromNodeDirFileGroupDo
	Not(conds ...gen.Condition) IPromNodeDirFileGroupDo
	Or(conds ...gen.Condition) IPromNodeDirFileGroupDo
	Select(conds ...field.Expr) IPromNodeDirFileGroupDo
	Where(conds ...gen.Condition) IPromNodeDirFileGroupDo
	Order(conds ...field.Expr) IPromNodeDirFileGroupDo
	Distinct(cols ...field.Expr) IPromNodeDirFileGroupDo
	Omit(cols ...field.Expr) IPromNodeDirFileGroupDo
	Join(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupDo
	Group(cols ...field.Expr) IPromNodeDirFileGroupDo
	Having(conds ...gen.Condition) IPromNodeDirFileGroupDo
	Limit(limit int) IPromNodeDirFileGroupDo
	Offset(offset int) IPromNodeDirFileGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromNodeDirFileGroupDo
	Unscoped() IPromNodeDirFileGroupDo
	Create(values ...*model.PromNodeDirFileGroup) error
	CreateInBatches(values []*model.PromNodeDirFileGroup, batchSize int) error
	Save(values ...*model.PromNodeDirFileGroup) error
	First() (*model.PromNodeDirFileGroup, error)
	Take() (*model.PromNodeDirFileGroup, error)
	Last() (*model.PromNodeDirFileGroup, error)
	Find() ([]*model.PromNodeDirFileGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromNodeDirFileGroup, err error)
	FindInBatches(result *[]*model.PromNodeDirFileGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromNodeDirFileGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromNodeDirFileGroupDo
	Assign(attrs ...field.AssignExpr) IPromNodeDirFileGroupDo
	Joins(fields ...field.RelationField) IPromNodeDirFileGroupDo
	Preload(fields ...field.RelationField) IPromNodeDirFileGroupDo
	FirstOrInit() (*model.PromNodeDirFileGroup, error)
	FirstOrCreate() (*model.PromNodeDirFileGroup, error)
	FindByPage(offset int, limit int) (result []*model.PromNodeDirFileGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromNodeDirFileGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	WhereID(ctx context.Context, id uint) (result *model.PromNodeDirFileGroup, err error)
}

// select * from @@table where id = @id
func (p promNodeDirFileGroupDo) WhereID(ctx context.Context, id uint) (result *model.PromNodeDirFileGroup, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("select * from prom_node_dir_file_groups where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p promNodeDirFileGroupDo) Debug() IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Debug())
}

func (p promNodeDirFileGroupDo) WithContext(ctx context.Context) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promNodeDirFileGroupDo) ReadDB() IPromNodeDirFileGroupDo {
	return p.Clauses(dbresolver.Read)
}

func (p promNodeDirFileGroupDo) WriteDB() IPromNodeDirFileGroupDo {
	return p.Clauses(dbresolver.Write)
}

func (p promNodeDirFileGroupDo) Session(config *gorm.Session) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Session(config))
}

func (p promNodeDirFileGroupDo) Clauses(conds ...clause.Expression) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promNodeDirFileGroupDo) Returning(value interface{}, columns ...string) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promNodeDirFileGroupDo) Not(conds ...gen.Condition) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promNodeDirFileGroupDo) Or(conds ...gen.Condition) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promNodeDirFileGroupDo) Select(conds ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promNodeDirFileGroupDo) Where(conds ...gen.Condition) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promNodeDirFileGroupDo) Order(conds ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promNodeDirFileGroupDo) Distinct(cols ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promNodeDirFileGroupDo) Omit(cols ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promNodeDirFileGroupDo) Join(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promNodeDirFileGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promNodeDirFileGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promNodeDirFileGroupDo) Group(cols ...field.Expr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promNodeDirFileGroupDo) Having(conds ...gen.Condition) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promNodeDirFileGroupDo) Limit(limit int) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promNodeDirFileGroupDo) Offset(offset int) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promNodeDirFileGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promNodeDirFileGroupDo) Unscoped() IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promNodeDirFileGroupDo) Create(values ...*model.PromNodeDirFileGroup) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promNodeDirFileGroupDo) CreateInBatches(values []*model.PromNodeDirFileGroup, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promNodeDirFileGroupDo) Save(values ...*model.PromNodeDirFileGroup) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promNodeDirFileGroupDo) First() (*model.PromNodeDirFileGroup, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroup), nil
	}
}

func (p promNodeDirFileGroupDo) Take() (*model.PromNodeDirFileGroup, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroup), nil
	}
}

func (p promNodeDirFileGroupDo) Last() (*model.PromNodeDirFileGroup, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroup), nil
	}
}

func (p promNodeDirFileGroupDo) Find() ([]*model.PromNodeDirFileGroup, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromNodeDirFileGroup), err
}

func (p promNodeDirFileGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromNodeDirFileGroup, err error) {
	buf := make([]*model.PromNodeDirFileGroup, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promNodeDirFileGroupDo) FindInBatches(result *[]*model.PromNodeDirFileGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promNodeDirFileGroupDo) Attrs(attrs ...field.AssignExpr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promNodeDirFileGroupDo) Assign(attrs ...field.AssignExpr) IPromNodeDirFileGroupDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promNodeDirFileGroupDo) Joins(fields ...field.RelationField) IPromNodeDirFileGroupDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promNodeDirFileGroupDo) Preload(fields ...field.RelationField) IPromNodeDirFileGroupDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promNodeDirFileGroupDo) FirstOrInit() (*model.PromNodeDirFileGroup, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroup), nil
	}
}

func (p promNodeDirFileGroupDo) FirstOrCreate() (*model.PromNodeDirFileGroup, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroup), nil
	}
}

func (p promNodeDirFileGroupDo) FindByPage(offset int, limit int) (result []*model.PromNodeDirFileGroup, count int64, err error) {
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

func (p promNodeDirFileGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promNodeDirFileGroupDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promNodeDirFileGroupDo) Delete(models ...*model.PromNodeDirFileGroup) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promNodeDirFileGroupDo) withDO(do gen.Dao) *promNodeDirFileGroupDo {
	p.DO = *do.(*gen.DO)
	return p
}
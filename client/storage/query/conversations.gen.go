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

	"dp_client/storage/model"
)

func newConversation(db *gorm.DB, opts ...gen.DOOption) conversation {
	_conversation := conversation{}

	_conversation.conversationDo.UseDB(db, opts...)
	_conversation.conversationDo.UseModel(&model.Conversation{})

	tableName := _conversation.conversationDo.TableName()
	_conversation.ALL = field.NewAsterisk(tableName)
	_conversation.AgentId = field.NewInt64(tableName, "agent_id")
	_conversation.ConversationUId = field.NewString(tableName, "conversation_uid")
	_conversation.ID = field.NewUint(tableName, "id")
	_conversation.CreatedAt = field.NewTime(tableName, "created_at")
	_conversation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_conversation.DeletedAt = field.NewField(tableName, "deleted_at")

	_conversation.fillFieldMap()

	return _conversation
}

type conversation struct {
	conversationDo

	ALL             field.Asterisk
	AgentId         field.Int64
	ConversationUId field.String
	ID              field.Uint
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (c conversation) Table(newTableName string) *conversation {
	c.conversationDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c conversation) As(alias string) *conversation {
	c.conversationDo.DO = *(c.conversationDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *conversation) updateTableName(table string) *conversation {
	c.ALL = field.NewAsterisk(table)
	c.AgentId = field.NewInt64(table, "agent_id")
	c.ConversationUId = field.NewString(table, "conversation_uid")
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *conversation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *conversation) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["agent_id"] = c.AgentId
	c.fieldMap["conversation_uid"] = c.ConversationUId
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c conversation) clone(db *gorm.DB) conversation {
	c.conversationDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c conversation) replaceDB(db *gorm.DB) conversation {
	c.conversationDo.ReplaceDB(db)
	return c
}

type conversationDo struct{ gen.DO }

type IConversationDo interface {
	gen.SubQuery
	Debug() IConversationDo
	WithContext(ctx context.Context) IConversationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IConversationDo
	WriteDB() IConversationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IConversationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IConversationDo
	Not(conds ...gen.Condition) IConversationDo
	Or(conds ...gen.Condition) IConversationDo
	Select(conds ...field.Expr) IConversationDo
	Where(conds ...gen.Condition) IConversationDo
	Order(conds ...field.Expr) IConversationDo
	Distinct(cols ...field.Expr) IConversationDo
	Omit(cols ...field.Expr) IConversationDo
	Join(table schema.Tabler, on ...field.Expr) IConversationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IConversationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IConversationDo
	Group(cols ...field.Expr) IConversationDo
	Having(conds ...gen.Condition) IConversationDo
	Limit(limit int) IConversationDo
	Offset(offset int) IConversationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IConversationDo
	Unscoped() IConversationDo
	Create(values ...*model.Conversation) error
	CreateInBatches(values []*model.Conversation, batchSize int) error
	Save(values ...*model.Conversation) error
	First() (*model.Conversation, error)
	Take() (*model.Conversation, error)
	Last() (*model.Conversation, error)
	Find() ([]*model.Conversation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Conversation, err error)
	FindInBatches(result *[]*model.Conversation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Conversation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IConversationDo
	Assign(attrs ...field.AssignExpr) IConversationDo
	Joins(fields ...field.RelationField) IConversationDo
	Preload(fields ...field.RelationField) IConversationDo
	FirstOrInit() (*model.Conversation, error)
	FirstOrCreate() (*model.Conversation, error)
	FindByPage(offset int, limit int) (result []*model.Conversation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IConversationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Conversation, err error)
}

// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (c conversationDo) FilterWithNameAndRole(name string, role string) (result []model.Conversation, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM conversations WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c conversationDo) Debug() IConversationDo {
	return c.withDO(c.DO.Debug())
}

func (c conversationDo) WithContext(ctx context.Context) IConversationDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c conversationDo) ReadDB() IConversationDo {
	return c.Clauses(dbresolver.Read)
}

func (c conversationDo) WriteDB() IConversationDo {
	return c.Clauses(dbresolver.Write)
}

func (c conversationDo) Session(config *gorm.Session) IConversationDo {
	return c.withDO(c.DO.Session(config))
}

func (c conversationDo) Clauses(conds ...clause.Expression) IConversationDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c conversationDo) Returning(value interface{}, columns ...string) IConversationDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c conversationDo) Not(conds ...gen.Condition) IConversationDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c conversationDo) Or(conds ...gen.Condition) IConversationDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c conversationDo) Select(conds ...field.Expr) IConversationDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c conversationDo) Where(conds ...gen.Condition) IConversationDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c conversationDo) Order(conds ...field.Expr) IConversationDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c conversationDo) Distinct(cols ...field.Expr) IConversationDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c conversationDo) Omit(cols ...field.Expr) IConversationDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c conversationDo) Join(table schema.Tabler, on ...field.Expr) IConversationDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c conversationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IConversationDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c conversationDo) RightJoin(table schema.Tabler, on ...field.Expr) IConversationDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c conversationDo) Group(cols ...field.Expr) IConversationDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c conversationDo) Having(conds ...gen.Condition) IConversationDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c conversationDo) Limit(limit int) IConversationDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c conversationDo) Offset(offset int) IConversationDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c conversationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IConversationDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c conversationDo) Unscoped() IConversationDo {
	return c.withDO(c.DO.Unscoped())
}

func (c conversationDo) Create(values ...*model.Conversation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c conversationDo) CreateInBatches(values []*model.Conversation, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c conversationDo) Save(values ...*model.Conversation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c conversationDo) First() (*model.Conversation, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Conversation), nil
	}
}

func (c conversationDo) Take() (*model.Conversation, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Conversation), nil
	}
}

func (c conversationDo) Last() (*model.Conversation, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Conversation), nil
	}
}

func (c conversationDo) Find() ([]*model.Conversation, error) {
	result, err := c.DO.Find()
	return result.([]*model.Conversation), err
}

func (c conversationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Conversation, err error) {
	buf := make([]*model.Conversation, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c conversationDo) FindInBatches(result *[]*model.Conversation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c conversationDo) Attrs(attrs ...field.AssignExpr) IConversationDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c conversationDo) Assign(attrs ...field.AssignExpr) IConversationDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c conversationDo) Joins(fields ...field.RelationField) IConversationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c conversationDo) Preload(fields ...field.RelationField) IConversationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c conversationDo) FirstOrInit() (*model.Conversation, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Conversation), nil
	}
}

func (c conversationDo) FirstOrCreate() (*model.Conversation, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Conversation), nil
	}
}

func (c conversationDo) FindByPage(offset int, limit int) (result []*model.Conversation, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c conversationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c conversationDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c conversationDo) Delete(models ...*model.Conversation) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *conversationDo) withDO(do gen.Dao) *conversationDo {
	c.DO = *do.(*gen.DO)
	return c
}

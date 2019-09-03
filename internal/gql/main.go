package gql

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
	models "github.com/wtlin1228/go-gql-server/internal/gql/models"
)

// region    ************************** generated!.gotpl **************************

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CreateUser func(childComplexity int, input models.UserInput) int
		DeleteUser func(childComplexity int, id string) int
		UpdateUser func(childComplexity int, id string, input models.UserInput) int
	}

	Query struct {
		Users func(childComplexity int, id *string) int
	}

	User struct {
		CreatedAt   func(childComplexity int) int
		Description func(childComplexity int) int
		Email       func(childComplexity int) int
		FirstName   func(childComplexity int) int
		ID          func(childComplexity int) int
		LastName    func(childComplexity int) int
		Location    func(childComplexity int) int
		Name        func(childComplexity int) int
		NickName    func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
		UserID      func(childComplexity int) int
	}

	Users struct {
		Count func(childComplexity int) int
		List  func(childComplexity int) int
	}
}

type MutationResolver interface {
	CreateUser(ctx context.Context, input models.UserInput) (*models.User, error)
	UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
}
type QueryResolver interface {
	Users(ctx context.Context, id *string) (*models.Users, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createUser":
		if e.complexity.Mutation.CreateUser == nil {
			break
		}

		args, err := ec.field_Mutation_createUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateUser(childComplexity, args["input"].(models.UserInput)), true

	case "Mutation.deleteUser":
		if e.complexity.Mutation.DeleteUser == nil {
			break
		}

		args, err := ec.field_Mutation_deleteUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteUser(childComplexity, args["id"].(string)), true

	case "Mutation.updateUser":
		if e.complexity.Mutation.UpdateUser == nil {
			break
		}

		args, err := ec.field_Mutation_updateUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateUser(childComplexity, args["id"].(string), args["input"].(models.UserInput)), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		args, err := ec.field_Query_users_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Users(childComplexity, args["id"].(*string)), true

	case "User.createdAt":
		if e.complexity.User.CreatedAt == nil {
			break
		}

		return e.complexity.User.CreatedAt(childComplexity), true

	case "User.description":
		if e.complexity.User.Description == nil {
			break
		}

		return e.complexity.User.Description(childComplexity), true

	case "User.email":
		if e.complexity.User.Email == nil {
			break
		}

		return e.complexity.User.Email(childComplexity), true

	case "User.firstName":
		if e.complexity.User.FirstName == nil {
			break
		}

		return e.complexity.User.FirstName(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.lastName":
		if e.complexity.User.LastName == nil {
			break
		}

		return e.complexity.User.LastName(childComplexity), true

	case "User.location":
		if e.complexity.User.Location == nil {
			break
		}

		return e.complexity.User.Location(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "User.nickName":
		if e.complexity.User.NickName == nil {
			break
		}

		return e.complexity.User.NickName(childComplexity), true

	case "User.updatedAt":
		if e.complexity.User.UpdatedAt == nil {
			break
		}

		return e.complexity.User.UpdatedAt(childComplexity), true

	case "User.userId":
		if e.complexity.User.UserID == nil {
			break
		}

		return e.complexity.User.UserID(childComplexity), true

	case "Users.count":
		if e.complexity.Users.Count == nil {
			break
		}

		return e.complexity.Users.Count(childComplexity), true

	case "Users.list":
		if e.complexity.Users.List == nil {
			break
		}

		return e.complexity.Users.List(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Mutation(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "internal/gql/schemas/schema.graphql", Input: `scalar Time
# Types
type User {
  id: ID!
  email: String!
  userId: String
  name: String
  firstName: String
  lastName: String
  nickName: String
  description: String
  location: String
  createdAt: Time!
  updatedAt: Time
}

# Input Types
input UserInput {
  email: String
  userId: String
  displayName: String
  name: String
  firstName: String
  lastName: String
  nickName: String
  description: String
  location: String
}

# List Types
type Users {
  count: Int # You want to return count for a grid for example
  list: [User!]! # that is why we need to specify the users object this way
}

# Define mutations here
type Mutation {
  createUser(input: UserInput!): User!
  updateUser(id: ID!, input: UserInput!): User!
  deleteUser(id: ID!): Boolean!
}

# Define queries here
type Query {
  users(id: ID): Users!
}`},
)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_createUser_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 models.UserInput
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = ec.unmarshalNUserInput2githubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUserInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteUser_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updateUser_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 models.UserInput
	if tmp, ok := rawArgs["input"]; ok {
		arg1, err = ec.unmarshalNUserInput2githubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUserInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_users_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalOID2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_createUser(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Mutation",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createUser_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateUser(rctx, args["input"].(models.UserInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*models.User)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNUser2ᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_updateUser(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Mutation",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_updateUser_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateUser(rctx, args["id"].(string), args["input"].(models.UserInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*models.User)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNUser2ᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_deleteUser(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Mutation",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_deleteUser_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteUser(rctx, args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_users(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_users_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Users(rctx, args["id"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*models.Users)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNUsers2ᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUsers(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query___type_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Query",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _User_email(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Email, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _User_userId(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_name(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_firstName(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.FirstName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_lastName(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.LastName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_nickName(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NickName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_description(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_location(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Location, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _User_createdAt(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) _User_updatedAt(ctx context.Context, field graphql.CollectedField, obj *models.User) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "User",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UpdatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) _Users_count(ctx context.Context, field graphql.CollectedField, obj *models.Users) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Users",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Count, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOInt2ᚖint(ctx, field.Selections, res)
}

func (ec *executionContext) _Users_list(ctx context.Context, field graphql.CollectedField, obj *models.Users) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "Users",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.List, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*models.User)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNUser2ᚕᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__DirectiveLocation2ᚕstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__TypeKind2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
		ec.Tracer.EndFieldExecution(ctx)
	}()
	rctx := &graphql.ResolverContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputUserInput(ctx context.Context, obj interface{}) (models.UserInput, error) {
	var it models.UserInput
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "email":
			var err error
			it.Email, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "userId":
			var err error
			it.UserID, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "displayName":
			var err error
			it.DisplayName, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "name":
			var err error
			it.Name, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "firstName":
			var err error
			it.FirstName, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "lastName":
			var err error
			it.LastName, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "nickName":
			var err error
			it.NickName, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "description":
			var err error
			it.Description, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "location":
			var err error
			it.Location, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, mutationImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createUser":
			out.Values[i] = ec._Mutation_createUser(ctx, field)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updateUser":
			out.Values[i] = ec._Mutation_updateUser(ctx, field)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deleteUser":
			out.Values[i] = ec._Mutation_deleteUser(ctx, field)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "users":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_users(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var userImplementors = []string{"User"}

func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *models.User) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, userImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "id":
			out.Values[i] = ec._User_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "email":
			out.Values[i] = ec._User_email(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "userId":
			out.Values[i] = ec._User_userId(ctx, field, obj)
		case "name":
			out.Values[i] = ec._User_name(ctx, field, obj)
		case "firstName":
			out.Values[i] = ec._User_firstName(ctx, field, obj)
		case "lastName":
			out.Values[i] = ec._User_lastName(ctx, field, obj)
		case "nickName":
			out.Values[i] = ec._User_nickName(ctx, field, obj)
		case "description":
			out.Values[i] = ec._User_description(ctx, field, obj)
		case "location":
			out.Values[i] = ec._User_location(ctx, field, obj)
		case "createdAt":
			out.Values[i] = ec._User_createdAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updatedAt":
			out.Values[i] = ec._User_updatedAt(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var usersImplementors = []string{"Users"}

func (ec *executionContext) _Users(ctx context.Context, sel ast.SelectionSet, obj *models.Users) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, usersImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Users")
		case "count":
			out.Values[i] = ec._Users_count(ctx, field, obj)
		case "list":
			out.Values[i] = ec._Users_list(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.RequestContext, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalNBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	res := graphql.MarshalBoolean(v)
	if res == graphql.Null {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNID2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalID(v)
}

func (ec *executionContext) marshalNID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalID(v)
	if res == graphql.Null {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}

func (ec *executionContext) marshalNTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := graphql.MarshalTime(v)
	if res == graphql.Null {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNUser2githubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx context.Context, sel ast.SelectionSet, v models.User) graphql.Marshaler {
	return ec._User(ctx, sel, &v)
}

func (ec *executionContext) marshalNUser2ᚕᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx context.Context, sel ast.SelectionSet, v []*models.User) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNUser2ᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalNUser2ᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUser(ctx context.Context, sel ast.SelectionSet, v *models.User) graphql.Marshaler {
	if v == nil {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
}

func (ec *executionContext) unmarshalNUserInput2githubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUserInput(ctx context.Context, v interface{}) (models.UserInput, error) {
	return ec.unmarshalInputUserInput(ctx, v)
}

func (ec *executionContext) marshalNUsers2githubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUsers(ctx context.Context, sel ast.SelectionSet, v models.Users) graphql.Marshaler {
	return ec._Users(ctx, sel, &v)
}

func (ec *executionContext) marshalNUsers2ᚖgithubᚗcomᚋwtlin1228ᚋgoᚑgqlᚑserverᚋinternalᚋgqlᚋmodelsᚋgeneratedᚐUsers(ctx context.Context, sel ast.SelectionSet, v *models.Users) graphql.Marshaler {
	if v == nil {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Users(ctx, sel, v)
}

func (ec *executionContext) marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v introspection.Directive) graphql.Marshaler {
	return ec.___Directive(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v []introspection.Directive) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalN__DirectiveLocation2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__DirectiveLocation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalN__DirectiveLocation2ᚕstring(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalN__DirectiveLocation2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalN__DirectiveLocation2ᚕstring(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__DirectiveLocation2string(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v introspection.EnumValue) graphql.Marshaler {
	return ec.___EnumValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v introspection.Field) graphql.Marshaler {
	return ec.___Field(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v introspection.InputValue) graphql.Marshaler {
	return ec.___InputValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func (ec *executionContext) unmarshalN__TypeKind2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__TypeKind2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
}

func (ec *executionContext) unmarshalOBoolean2ᚖbool(ctx context.Context, v interface{}) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOBoolean2bool(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOBoolean2ᚖbool(ctx context.Context, sel ast.SelectionSet, v *bool) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOBoolean2bool(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOID2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalID(v)
}

func (ec *executionContext) marshalOID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalID(v)
}

func (ec *executionContext) unmarshalOID2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOID2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOID2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOID2string(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOInt2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalInt(v)
}

func (ec *executionContext) marshalOInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	return graphql.MarshalInt(v)
}

func (ec *executionContext) unmarshalOInt2ᚖint(ctx context.Context, v interface{}) (*int, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOInt2int(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOInt2ᚖint(ctx context.Context, sel ast.SelectionSet, v *int) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOInt2int(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalOString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOString2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOString2string(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}

func (ec *executionContext) marshalOTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	return graphql.MarshalTime(v)
}

func (ec *executionContext) unmarshalOTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOTime2timeᚐTime(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOTime2timeᚐTime(ctx, sel, *v)
}

func (ec *executionContext) marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v []introspection.EnumValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v []introspection.Field) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Schema2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v introspection.Schema) graphql.Marshaler {
	return ec.___Schema(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v *introspection.Schema) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, sel, v)
}

func (ec *executionContext) marshalO__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************

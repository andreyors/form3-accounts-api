我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

Form3 Take Home Exercise
==

# Technical decisions
- I've started the task discovery from [API doc](https://api-docs.form3.tech/api.html#organisation-accounts)
- What's out of scope by the task def:
    - We don't need to provide an authorisation to access Fake API
    - We cannot use swagger autogen tool
    - We need to use low-level HTTP communication (net/http)
- We need to introduce `Create`, `Fetch`, `List` and `Delete` operations for accounts
- We need to keep things aligned, so we will introduce the sameish naming convention for client library methods
- The API provides one extra endpoint `Patch`, we don't need to introduce it
- I've started provided docker env
- I manually checked if all required endpoints are here and can be used
    - `Delete` does not behave as expected, it basically returns always 200 even for non-existent version
    - `Patch` endpoint does not exist
- I've updated the docker tag to latest and checked again - nothing changed for previous bullet point
- I personally prefer start from data model
    - [There](https://api-docs.form3.tech/assets/form3-swagger.yaml) I found documentation
    - There are some small inconsistencies on required fields
- I've added `/cmd/form3/models` folder for all models
- File `models/account.go` contains a detailed defintion for `account`, reversed engineered from YAML. It will be useful
for `Create` (marshall/unmarshall) and `Fetch` (unmarshall) operations
    - The `account` definition in YAML is quite big and we need to limit the complexity
    - I took only relevant for the first version set of fields
- File `models/accounts.go` contains an account collection definition for `List` operation, actually it is only an array
of `models/account.go` entities
- For `List` endpoint we need to introduce only pagination, so we added `models/pagination.go` for pager definition - it
contains two fields - `Number` (default 0), `Size` (default 100)
- I played with Postman, and discover [HATEOAS media links](https://en.wikipedia.org/wiki/HATEOAS), so I've added `models/links.go`
- From perspective of production readiness I can mention:
    - It should have all meaningful unit tests
    - It should have all meaningful e2e tests
    - It should be easily extensible
        - We need to keep SRP principle applied - so we try to keep logic separated
        - New features for entity management should be simple and straighforward
    - It should log all important steps
- The idea of pluggable API can be introduced via interfaces
    - So we added `cmd/form3/form3.go` which contains the API aggregator
```go
type Form3 struct {
  Account Account.ClientInterface
}

func New(host string, timeout time.Duration) *Form3 {
  return &Form3{
    Account: Account.NewClient(host, timeout),
  }
}

```
And the aggregator will be called in the unified way
```go
f3 := Form3.New(os.Getenv('FORM3_API'), time.Duration(5 * time.Second))

f3.Account.Create(...)
```
We can extend this aggregator in the future
```go
type Form3 struct {
  Account Account.ClientInterface
  Claim Claim.ClientInterface
  Mandate Mandate.ClientInterface
}

func New(host string, timeout time.Duration) *Form3 {
  return &Form3{
    Account: Account.NewClient(host, timeout),
    Claim: Claim.NewClient(host, timeout),
    Mandate: Mandate.NewClient(host, timeout),
  }
}

f3.Mandate.Create(...)
f3.Claim.List(...)
```
- We need to follow 12-factor app principles
    - We introduced automatic tests for docker image, check `CMD`
    - We inject `FORM3_API` via env
- We keep models simple, without any logic

# What is left/can be improved
- Implement [complex validation rules](https://api-docs.form3.tech/api.html#organisation-accounts-create) for all
countries (currently done only for GB, take a look in the `validation` folder and [docs](https://pkg.go.dev/gopkg.in/go-playground/validator.v9?tab=doc))
- Implement [IBAN](https://api-docs.form3.tech/api.html#organisation-accounts-account-number-generation) generation
- Implement [PATCH](https://api-docs.form3.tech/api.html#organisation-accounts-patch) operation
- OrganisationID must be injected into Form3
- Models shared between different domains
- Add CI/CD

## Tech stack
- Golang 1.14
- [UUID](https://github.com/google/uuid) from Google
- [validation.v9](https://github.com/go-playground/validator) library with tag-based validation
- testify/assert & testify/mock

## Setup
0. Install docker and docker-compose
1. Run `docker-compose up`
2. Press Ctrl-C after test review

## Project structure
```
\_ cmd/form3
    \_ account      # account client interface implementation
    \_ common       # common helpers
    \_ models       # shared models
    \_ test         # test helpers
    \_ validation   # validators
    form3           # API aggregator
\_ scripts
    \_ db           # db init scripts
    \_ test         # JSON fixture folder
Makefile            # useful shortcuts
```


package biz

type Query struct {
	Page int64
	Size int64
	Name string
}

type BasePolicy struct {
	BaseInfo
	Manifest string
}

type BaseInfo struct {
	ID          int64
	Name        string
	SystemShort string
	Environment string
}

type SpecPolicy struct {
	BaseInfo
	Answers        map[string]interface{}
	BaseYaml       string
	Manifest       string
	SpecTemplateID int64
}

type ImportPolicy struct {
	BaseInfo
	Args       map[string]interface{}
	TemplateID int64
	Manifest   string
}

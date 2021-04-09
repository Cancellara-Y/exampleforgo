package option


var defaultOption = templateApi{
	path :"localpath",
	name : "api",
}

type TemplateApi interface {
	DoRequest() error
 }

type templateApi struct{
	path string
	name string
}

type TemplateApiOption func(*templateApi)

func WithPath(path string) TemplateApiOption {
	return func(t *templateApi){
		t.path = path
	}
}

func WithName(name string) TemplateApiOption {
	return func(t *templateApi){
		t.name = name
	}
}

func NewTemplateApiOption(opts ...TemplateApiOption) TemplateApi{
	api := defaultOption
	for _, v:= range opts{
		v(&api)
	}
	return api
}

func (t templateApi)DoRequest() error{
	return nil
}
package action

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"html/template"
)

type AjaxAction struct {
	BtnId    string
	Url      string
	Method   string
	Title    string
	Data     AjaxData
	Handlers []context.Handler
}

func Ajax(url, title string, handler Handler) *AjaxAction {
	return &AjaxAction{
		Url:      url,
		Title:    title,
		Method:   "post",
		Data:     NewAjaxData(),
		Handlers: context.Handlers{handler.Wrap()},
	}
}

func (ajax *AjaxAction) SetData(data map[string]interface{}) *AjaxAction {
	ajax.Data = ajax.Data.Add(data)
	return ajax
}

func (ajax *AjaxAction) SetUrl(url string) *AjaxAction {
	ajax.Url = url
	return ajax
}

func (ajax *AjaxAction) SetMethod(method string) *AjaxAction {
	ajax.Method = method
	return ajax
}

func (ajax *AjaxAction) GetCallbacks() context.Node {
	return context.Node{
		Path:     ajax.Url,
		Method:   ajax.Method,
		Handlers: ajax.Handlers,
		Value:    map[string]interface{}{auth.ContextNodeNeedAuth: 1},
	}
}

func (ajax *AjaxAction) SetBtnId(btnId string) {
	ajax.BtnId = btnId
}

func (ajax *AjaxAction) Js() template.JS {
	return template.JS(`$('#` + ajax.BtnId + `').on('click', function (event) {
						$.ajax({
                            method: '` + ajax.Method + `',
                            url: "` + ajax.Url + `",
                            data: ` + ajax.Data.JSON() + `,
                            success: function (data) { 
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code === 0) {
                                    swal(data.msg, '', 'success');
                                } else {
                                    swal(data.msg, '', 'error');
                                }
                            }
                        });
            		});`)
}

func (ajax *AjaxAction) BtnAttribute() template.HTML { return template.HTML(``) }
func (ajax *AjaxAction) ExtContent() template.HTML   { return template.HTML(``) }

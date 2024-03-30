// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package errors_pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ppp3ppj/wywy/views/layout"
import "github.com/ppp3ppj/wywy/models"

func Error401(fromProtected bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-base-error-agent h-screen justify-center\"><center class=\"mt-24 m-auto\"><svg class=\"emoji-401\" enable-background=\"new 0 0 226 249.135\" height=\"249.135\" id=\"Layer_1\" overflow=\"visible\" version=\"1.1\" viewBox=\"0 0 226 249.135\" width=\"226\" xml:space=\"preserve\"><circle cx=\"113\" cy=\"113\" fill=\"#FFE585\" r=\"109\"></circle><line enable-background=\"new    \" fill=\"none\" opacity=\"0.29\" stroke=\"#6E6E96\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\" x1=\"88.866\" x2=\"136.866\" y1=\"245.135\" y2=\"245.135\"></line><line enable-background=\"new    \" fill=\"none\" opacity=\"0.17\" stroke=\"#6E6E96\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\" x1=\"154.732\" x2=\"168.732\" y1=\"245.135\" y2=\"245.135\"></line><line enable-background=\"new    \" fill=\"none\" opacity=\"0.17\" stroke=\"#6E6E96\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\" x1=\"69.732\" x2=\"58.732\" y1=\"245.135\" y2=\"245.135\"></line><circle cx=\"68.732\" cy=\"93\" fill=\"#6E6E96\" r=\"9\"></circle><path d=\"M115.568,5.947c-1.026,0-2.049,0.017-3.069,0.045  c54.425,1.551,98.069,46.155,98.069,100.955c0,55.781-45.219,101-101,101c-55.781,0-101-45.219-101-101  c0-8.786,1.124-17.309,3.232-25.436c-3.393,10.536-5.232,21.771-5.232,33.436c0,60.199,48.801,109,109,109s109-48.801,109-109  S175.768,5.947,115.568,5.947z\" enable-background=\"new    \" fill=\"#FF9900\" opacity=\"0.24\"></path><circle cx=\"156.398\" cy=\"93\" fill=\"#6E6E96\" r=\"9\"></circle><ellipse cx=\"67.732\" cy=\"140.894\" enable-background=\"new    \" fill=\"#FF0000\" opacity=\"0.18\" rx=\"17.372\" ry=\"8.106\"></ellipse><ellipse cx=\"154.88\" cy=\"140.894\" enable-background=\"new    \" fill=\"#FF0000\" opacity=\"0.18\" rx=\"17.371\" ry=\"8.106\"></ellipse><path d=\"M13,118.5C13,61.338,59.338,15,116.5,15c55.922,0,101.477,44.353,103.427,99.797  c0.044-1.261,0.073-2.525,0.073-3.797C220,50.802,171.199,2,111,2S2,50.802,2,111c0,50.111,33.818,92.318,79.876,105.06  C41.743,201.814,13,163.518,13,118.5z\" fill=\"#FFEFB5\"></path><circle cx=\"113\" cy=\"113\" fill=\"none\" r=\"109\" stroke=\"#6E6E96\" stroke-width=\"8\"></circle></svg><div class=\"mt-4\"><span class=\"text-6xl md:text-6xl lg:text-8xl font-bold block\"><span>4  0  1</span></span> <span class=\"text-2xl md:text-2xl lg:text-2xl block\">Status Unauthorized</span> <span class=\"text-xl md:text-xl lg:text-xl\">Please provide valid credentials.</span></div></center> <center class=\"mt-6\"><a hx-swap=\"transition:true\" href=\"/login\" class=\"text-xl md:text-xl lg:text-xl p-3 rounded-md hover:shadow-md\">Go Login Page </a></center></div><style>\n      body, html {\n        height: 100%;\n        margin: 0;\n        display: flex;\n        align-items: center;\n        justify-content: center;\n      }\n\n      .bg-base-error-agent {\n        width: 100%;\n        max-width: 400px; /* Set maximum width for content */\n        padding: 20px;\n        box-sizing: border-box;\n        text-align: center;\n      }\n\n      .emoji-401 {\n        position: relative;\n        animation: mymove 2.5s infinite;\n      }\n\n      @keyframes mymove {\n        33% {top: 0px;}\n        66% {top: 20px;}\n        100% {top: 0px}\n      }\n    }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ErorrIndex(
	u models.UserNav,
	fromProtected,
	isError bool,
	cmp templ.Component,
) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = cmp.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base(u, fromProtected, false, false).Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

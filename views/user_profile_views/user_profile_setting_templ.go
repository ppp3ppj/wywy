// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package user_profile_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ppp3ppj/wywy/views/layout"
import "github.com/ppp3ppj/wywy/models"

func UserProfileSetting() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"p-8 w-9/12\"><div hx-swap=\"transition:true\"><div class=\"bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col my-2\"><label class=\"block uppercase tracking-wide text-grey-darker text-2xl font-bold mb-2\">Profile Setting</label><div class=\"avatar\"><div class=\"w-24 rounded\" hx-get=\"/ppp-image-agent\"><img src=\"https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg\"></div></div><div class=\"-mx-3 md:flex mb-6\"><div class=\"md:w-1/2 px-3 mb-6 md:mb-0\"><label class=\"block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2\" for=\"grid-password\">Firstname</label> <label class=\"form-control\"><div class=\"label\"><span class=\"label-text\">What is your name?</span> <span class=\"label-text-alt\">Top Right label</span></div><input type=\"text\" placeholder=\"Firstname\" class=\"input input-bordered\"><div class=\"label\"><span class=\"label-text-alt\">Bottom Left label</span> <span class=\"label-text-alt\">Bottom Right label</span></div></label></div><div class=\"md:w-1/2 px-3\"><label class=\"block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2\" for=\"grid-password\">Lastname</label> <label class=\"form-control\"><div class=\"label\"><span class=\"label-text\">What is your name?</span> <span class=\"label-text-alt\">Top Right label</span></div><input type=\"text\" placeholder=\"Lastname\" class=\"input input-bordered\"><div class=\"label\"><span class=\"label-text-alt\">Bottom Left label</span> <span class=\"label-text-alt\">Bottom Right label</span></div></label></div></div><div class=\"-mx-3 md:flex mb-6\"><div class=\"md:w-full px-3\"><label class=\"block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2\" for=\"grid-password\">Password</label> <input class=\"appearance-none block w-full bg-grey-lighter text-grey-darker border border-grey-lighter rounded py-3 px-4 mb-3\" id=\"grid-password\" type=\"password\" placeholder=\"******************\"><p class=\"text-grey-dark text-xs italic\">Make it as long and as crazy as you'd like</p></div></div><div class=\"-mx-3 md:flex mb-2\"><div class=\"md:w-1/2 px-3 mb-6 md:mb-0\"><label class=\"block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2\" for=\"grid-city\">City</label> <input class=\"appearance-none block w-full bg-grey-lighter text-grey-darker border border-grey-lighter rounded py-3 px-4\" id=\"grid-city\" type=\"text\" placeholder=\"Albuquerque\"></div><div class=\"md:w-1/2 px-3\"><label class=\"block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2\" for=\"grid-state\">State</label><div class=\"relative\"><select class=\"block appearance-none w-full bg-grey-lighter border border-grey-lighter text-grey-darker py-3 px-4 pr-8 rounded\" id=\"grid-state\"><option>New Mexico</option> <option>Missouri</option> <option>Texas</option></select><div class=\"pointer-events-none absolute pin-y pin-r flex items-center px-2 text-grey-darker\"><svg class=\"h-4 w-4\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\"><path d=\"M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z\"></path></svg></div></div></div><div class=\"md:w-1/2 px-3\"><label class=\"block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2\" for=\"grid-zip\">Zip</label> <input class=\"appearance-none block w-full bg-grey-lighter text-grey-darker border border-grey-lighter rounded py-3 px-4\" id=\"grid-zip\" type=\"text\" placeholder=\"90210\"></div></div><button hx-swap=\"transition:true\" class=\"btn btn-neutral w-3/12\">Save</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func UserProfileSettingIndex(u models.UserNav, fromProtected bool, cmp templ.Component) templ.Component {
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
		templ_7745c5c3_Err = layout.Base(u, fromProtected, true, true).Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

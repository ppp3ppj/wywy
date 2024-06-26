// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package signup

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ppp3ppj/wywy/pub"

func Signup() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head hx-boost=\"true\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = pub.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>Signup-ppp</title></head><body id=\"admin-root\" class=\"bg-base-100\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = signup().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func signup() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-grey-lighter min-h-screen flex flex-col\"><div class=\"container max-w-sm mx-auto flex-1 flex flex-col items-center justify-center px-2\"><div class=\"bg-white px-6 py-8 rounded shadow-md text-black w-full\"><h1 class=\"mb-8 text-3xl text-center\">Sign up</h1><form hx-swap=\"transition:true\" hx-post=\"/signup\"><div class=\"mb-3\"><input id=\"firstname\" name=\"firstname\" type=\"text\" placeholder=\"Firstname\" class=\"input input-bordered input-primary w-full max-w-xs\"></div><div class=\"mb-3\"><input id=\"lastname\" name=\"lastname\" type=\"text\" placeholder=\"Lastname\" class=\"input input-bordered input-primary w-full max-w-xs\"></div><div class=\"mb-3\"><input id=\"username\" name=\"username\" type=\"text\" placeholder=\"Username\" class=\"input input-bordered input-primary w-full max-w-xs\"></div><div class=\"mb-3\"><input id=\"email\" name=\"email\" type=\"text\" placeholder=\"Email\" placeholder=\"Type here\" class=\"input input-bordered input-primary w-full max-w-xs p-4\"></div><div class=\"mb-3\"><input id=\"password\" name=\"password\" type=\"password\" placeholder=\"Password\" class=\"input input-bordered input-primary w-full max-w-xs\"></div><div class=\"mb-3\"><input id=\"confirm_password\" name=\"confirm_password\" type=\"password\" placeholder=\"Confirm Password\" class=\"input input-bordered input-primary w-full max-w-xs\"></div><button class=\"btn btn-active btn-accent w-full\">Create</button></form><div class=\"text-center text-sm text-grey-dark mt-4\">By signing up, you agree to the <a class=\"no-underline border-b border-grey-dark text-grey-dark\" href=\"#\">Terms of Service</a> and <a class=\"no-underline border-b border-grey-dark text-grey-dark\" href=\"#\">Privacy Policy</a></div></div><div class=\"text-grey-dark mt-6\">Already have an account? <a hx-swap=\"transition:true\" hx-get=\"/login\" hx-target=\"body\" hx-push-url=\"true\" class=\"no-underline border-b border-blue text-blue\">Log in</a>.</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

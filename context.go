package fasthttp

import (
	"context"
	"time"
)

type contextKey struct {
	name string
}

var (
	ctxKey = &contextKey{"context"}
)

func (ctx *RequestCtx) Context() context.Context {
	if c, ok := ctx.UserValue(ctxKey).(context.Context); ok {
		return c
	}
	return nil
}

func (ctx *RequestCtx) SetContext(c context.Context) {
	ctx.SetUserValue(ctxKey, c)
}

// Method: as is
// RequestURI: as is
func (ctx *RequestCtx) HostBytes() []byte {
	return ctx.Host()
}
func (ctx *RequestCtx) PathBytes() []byte {
	return ctx.Path()
}
func (ctx *RequestCtx) GetHeader(name string) []byte {
	return ctx.Request.Header.Peek(name)
}
func (ctx *RequestCtx) CookieBytes(name string) []byte {
	return ctx.Request.Header.Cookie(name)
}

func (ctx *RequestCtx) FormValueBytes(name string) []byte {
	return ctx.FormValue(name)
}

func (ctx *RequestCtx) FormValueString(name string) string {
	return string(ctx.FormValue(name))
}

// MultipartForm: as is

func (ctx *RequestCtx) ReadBody() []byte {
	return ctx.PostBody()
}

// SetStatusCode: as is
// SetContentType: as is
// SetContentTypeBytes: as is

func (ctx *RequestCtx) SetHeader(name, val string) {
	ctx.Response.Header.Set(name, val)
}
func (ctx *RequestCtx) SetHeaderBytes(name, val []byte) {
	ctx.Response.Header.SetBytesKV(name, val)
}

// SetBody: as is

//func (ctx *RequestCtx) SetCookie(ck *http.Cookie) {
//	cookie := &Cookie{
//		key:      []byte(ck.Name),
//		value:    []byte(ck.Value),
//		expire:   ck.Expires,
//		maxAge:   ck.MaxAge,
//		domain:   []byte(ck.Domain),
//		path:     []byte(ck.Path),
//		httpOnly: ck.HttpOnly,
//		secure:   ck.Secure,
//		sameSite: CookieSameSite(ck.SameSite),
//	}
//	ctx.Response.Header.SetCookie(cookie)
//}

func (ctx *RequestCtx) SetCookieBytes(key, val, domain, path []byte, expire time.Time, httpOnly, secure bool, sameSite int) {
	var maxAge int
	if !expire.IsZero() && expire.Before(time.Now()) {
		maxAge = -1
	}
	cookie := &Cookie{
		key:      key,
		value:    val,
		expire:   expire,
		maxAge:   maxAge,
		domain:   domain,
		path:     path,
		httpOnly: httpOnly,
		secure:   secure,
		sameSite: CookieSameSite(sameSite),
	}
	ctx.Response.Header.SetCookie(cookie)
}

func (ctx *RequestCtx) RedirectCode(uri string, code int) {
	ctx.Redirect(uri, code)
}

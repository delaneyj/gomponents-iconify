package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Secure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaSecure0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaSecure1" d="m 127.99999,239.96468 c 0,0 95.98506,-31.99503 95.98506,-111.98257"/><path id="galaSecure2" d="M 223.98505,127.98211 V 31.997059 c 0,0 -31.99502,-15.997511 -95.98506,-15.997511"/><path id="galaSecure3" d="m 128,239.96468 c 0,0 -95.985056,-31.99503 -95.985056,-111.98257"/><path id="galaSecure4" d="M 32.014944,127.98211 V 31.997059 c 0,0 31.995019,-15.997509 95.985056,-15.997509"/><path id="galaSecure5" d="M 191.99003,63.99208 C 128,111.9846 112.00249,175.97464 112.00249,175.97464 c 0,0 -15.997511,-19.0946 -31.995019,-31.99502"/></g>`),
		g.Group(children),
	)
}
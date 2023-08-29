package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinShake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWeixinShake0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 19L29 6L6 29l13 13l23-23Z"/><path stroke="#000" d="m16 29l3 3"/><path stroke="#fff" d="m30 42l12-12M6 18L18 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWeixinShake0)"/>`),
		g.Group(children),
	)
}
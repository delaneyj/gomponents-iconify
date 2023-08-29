package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinShake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWeixinShake0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M42 19L29 6L6 29l13 13l23-23Z"/><path d="m16 29l3 3m11 10l12-12M6 18L18 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWeixinShake0)"/>`),
		g.Group(children),
	)
}
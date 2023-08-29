package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gongfu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGongfu0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="22" cy="8" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 18h10v11h-7v14m25.182-25H26v10.86L40 43"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGongfu0)"/>`),
		g.Group(children),
	)
}
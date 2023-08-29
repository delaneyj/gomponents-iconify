package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTent0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M10 12L4 36h12"/><path fill="#fff" stroke="#fff" d="M38 12H10l6 24h28l-6-24Z"/><path stroke="#000" d="M12 18h27"/><path stroke="#fff" d="m10 12l3 12m25-12l3 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTent0)"/>`),
		g.Group(children),
	)
}
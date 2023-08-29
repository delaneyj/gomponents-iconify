package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdmiCableOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 4V.5h8V4M5 2.5h1m3 0h1M5.5 13v2m4-2v2m-7-10.5h10v5l-2 1v2h-6v-2l-2-1v-5Z"/>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Computer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="10" height="9" x="19" y="32" stroke="#000" stroke-linecap="round" stroke-linejoin="round"/><rect width="38" height="24" x="5" y="8" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M22 27H26"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M14 41L34 41"/></g>`),
		g.Group(children),
	)
}
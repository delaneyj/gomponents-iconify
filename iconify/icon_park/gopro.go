package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gopro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="13" height="34" x="5" y="7" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><rect width="19" height="22" x="24" y="13" fill="#2F88FF" stroke="#000" stroke-width="4" rx="3"/><circle cx="33.5" cy="24.5" r="3.5" fill="#43CCF8" stroke="#fff" stroke-width="4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 24H24"/><rect width="5" height="5" x="9" y="15" fill="#000" rx="2.5"/></g>`),
		g.Group(children),
	)
}
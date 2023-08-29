package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="38" height="28" x="5" y="13" fill="#2F88FF" stroke="#000" rx="2"/><circle cx="18" cy="28" r="6" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M30 24L36 24"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M30 32L36 32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M7 13L25 4"/></g>`),
		g.Group(children),
	)
}
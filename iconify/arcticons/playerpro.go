package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playerpro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="18.5"/><circle cx="24" cy="24" r="14.705"/><path d="M30.494 24L21.3 29.308V18.692L30.494 24Z"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7.635" cy="7.635" r="2.135"/><path d="M9.113 9.174L6.096 6.156"/><circle cx="7.635" cy="40.365" r="2.135"/><path d="m9.174 38.887l-3.018 3.017"/><circle cx="40.365" cy="7.635" r="2.135"/><path d="m38.887 9.174l3.017-3.018"/><circle cx="40.365" cy="40.365" r="2.135"/><path d="m38.826 38.887l3.018 3.017"/></g>`),
		g.Group(children),
	)
}
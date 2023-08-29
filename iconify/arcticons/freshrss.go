package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freshrss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5A21.5 21.5 0 1 0 2.5 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 39.5A15.5 15.5 0 1 0 8.5 24"/><circle cx="24" cy="24" r="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Omniclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M25.5 24a1.5 1.5 0 1 1-1.5-1.5a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".959" d="m25.4 23.476l12.58-3.98m-27.305 17.82L22.9 25.01"/>`),
		g.Group(children),
	)
}
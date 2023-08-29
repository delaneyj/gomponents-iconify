package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 19H3l9-15m8.615 11.171h.015m-1.115-3.4h.015m-1.815-3.1h.015m-2.315-2.7h.015"/>`),
		g.Group(children),
	)
}
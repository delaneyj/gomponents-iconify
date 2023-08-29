package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5A14.82 14.82 0 0 0 9.18 19.32h0v1.08c.6 8.12 7.34 14.65 14.82 23.1c7.81-8.82 14.82-15.5 14.82-24.18h0A14.82 14.82 0 0 0 24 4.5Zm0 7.7a7.13 7.13 0 1 1-7.13 7.12A7.13 7.13 0 0 1 24 12.2Z"/>`),
		g.Group(children),
	)
}
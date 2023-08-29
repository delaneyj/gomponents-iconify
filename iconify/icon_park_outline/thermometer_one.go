package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24 44a9 9 0 0 0 4-17.064V10c0-2 0-6-4-6s-4 4-4 6v16.936A9 9 0 0 0 24 44Z"/>`),
		g.Group(children),
	)
}
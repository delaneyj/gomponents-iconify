package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureFahrenheit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8a2 2 0 1 0 4 0a2 2 0 1 0-4 0m9 4h5m2-6h-6a1 1 0 0 0-1 1v11"/>`),
		g.Group(children),
	)
}
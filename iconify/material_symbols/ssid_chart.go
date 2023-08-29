package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SsidChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-6.2-6L3 17v-2.45l3-2.15l6.125 5.95L16.3 15H21v2h-4l-5 4Zm0-9L7.625 7.625L3 11V8.525L7.825 5L12.2 9.375L21 3v2.475L12 12Z"/>`),
		g.Group(children),
	)
}
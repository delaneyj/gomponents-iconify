package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProcessChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.3 18.45l-1.8-.9l6-12l1.8.9l-6 12Zm6.6 0l-1.8-.9l6-12l1.8.9l-6 12Zm6.6 0l-1.8-.9l6-12l1.8.9l-6 12Z"/>`),
		g.Group(children),
	)
}
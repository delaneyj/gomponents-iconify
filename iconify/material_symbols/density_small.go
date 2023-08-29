package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DensitySmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-2h18v2H3Zm0-6v-2h18v2H3Zm0-6V8h18v2H3Zm0-6V2h18v2H3Z"/>`),
		g.Group(children),
	)
}
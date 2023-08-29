package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureNegOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14H3v-2h7v2Zm5.75 5V8.05l-2.3 1.65l-1.15-1.75L16.4 5H18v14h-2.25Z"/>`),
		g.Group(children),
	)
}
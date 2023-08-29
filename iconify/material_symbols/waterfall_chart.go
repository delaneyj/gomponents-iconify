package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-7h3v7H3Zm4-6v-4h3v4H7Zm3-5V5h3v4h-3Zm4-2V4h3v3h-3Zm4 13V4h3v16h-3Z"/>`),
		g.Group(children),
	)
}
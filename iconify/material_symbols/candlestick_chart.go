package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandlestickChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20v-2H5V6h2V4h2v2h2v12H9v2H7Zm8 0v-5h-2V8h2V4h2v4h2v7h-2v5h-2Z"/>`),
		g.Group(children),
	)
}
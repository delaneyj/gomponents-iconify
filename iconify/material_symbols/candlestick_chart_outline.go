package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandlestickChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20v-2H5V6h2V4h2v2h2v12H9v2H7Zm0-4h2V8H7v8Zm8 4v-5h-2V8h2V4h2v4h2v7h-2v5h-2Zm0-7h2v-3h-2v3Zm-7-1Zm8-.5Z"/>`),
		g.Group(children),
	)
}
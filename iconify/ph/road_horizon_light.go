package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadHorizonLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.94 197.23a6 6 0 0 1-8.17-2.29L156.49 70H134v10a6 6 0 0 1-12 0V70H99.51L29.23 194.94a6 6 0 1 1-10.46-5.88L85.74 70H24a6 6 0 0 1 0-12h208a6 6 0 0 1 0 12h-61.74l67 119.06a6 6 0 0 1-2.32 8.17ZM128 114a6 6 0 0 0-6 6v16a6 6 0 0 0 12 0v-16a6 6 0 0 0-6-6Zm0 56a6 6 0 0 0-6 6v16a6 6 0 0 0 12 0v-16a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationChartThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44h-84V24a4 4 0 0 0-8 0v20H40a12 12 0 0 0-12 12v120a12 12 0 0 0 12 12h47.68l-26.8 33.5a4 4 0 1 0 6.24 5l30.8-38.5h60.16l30.8 38.5a4 4 0 0 0 6.24-5l-26.8-33.5H216a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm4 132a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4V56a4 4 0 0 1 4-4h176a4 4 0 0 1 4 4Zm-120-56v24a4 4 0 0 1-8 0v-24a4 4 0 0 1 8 0Zm32-16v40a4 4 0 0 1-8 0v-40a4 4 0 0 1 8 0Zm32-16v56a4 4 0 0 1-8 0V88a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}
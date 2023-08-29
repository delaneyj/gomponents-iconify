package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m240.49 119.51l-96-96A12 12 0 0 0 124 32v36h-4a12 12 0 0 0-12 12v96a12 12 0 0 0 12 12h4v36a12 12 0 0 0 20.49 8.49l96-96a12 12 0 0 0 0-16.98ZM148 195v-19a12 12 0 0 0-12-12h-4V92h4a12 12 0 0 0 12-12V61l67 67ZM52 80v96a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0Zm40 0v96a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}
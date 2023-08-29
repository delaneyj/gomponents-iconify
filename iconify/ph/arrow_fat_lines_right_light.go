package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m236.24 123.76l-96-96A6 6 0 0 0 130 32v42h-26a6 6 0 0 0-6 6v96a6 6 0 0 0 6 6h26v42a6 6 0 0 0 10.24 4.24l96-96a6 6 0 0 0 0-8.48ZM142 209.51V176a6 6 0 0 0-6-6h-26V86h26a6 6 0 0 0 6-6V46.49L223.51 128ZM46 80v96a6 6 0 0 1-12 0V80a6 6 0 0 1 12 0Zm32 0v96a6 6 0 0 1-12 0V80a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}
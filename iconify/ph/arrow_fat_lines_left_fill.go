package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 80v96a8 8 0 0 1-8 8h-24v40a8 8 0 0 1-13.66 5.66l-96-96a8 8 0 0 1 0-11.32l96-96A8 8 0 0 1 128 32v40h24a8 8 0 0 1 8 8Zm24-8a8 8 0 0 0-8 8v96a8 8 0 0 0 16 0V80a8 8 0 0 0-8-8Zm32 0a8 8 0 0 0-8 8v96a8 8 0 0 0 16 0V80a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
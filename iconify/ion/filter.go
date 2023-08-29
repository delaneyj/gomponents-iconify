package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 168H40a24 24 0 0 1 0-48h432a24 24 0 0 1 0 48Zm-80 112H120a24 24 0 0 1 0-48h272a24 24 0 0 1 0 48Zm-96 112h-80a24 24 0 0 1 0-48h80a24 24 0 0 1 0 48Z"/>`),
		g.Group(children),
	)
}
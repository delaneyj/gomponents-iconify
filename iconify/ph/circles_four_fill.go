package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 80a40 40 0 1 1-40-40a40 40 0 0 1 40 40Zm56 40a40 40 0 1 0-40-40a40 40 0 0 0 40 40Zm-96 16a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm96 0a40 40 0 1 0 40 40a40 40 0 0 0-40-40Z"/>`),
		g.Group(children),
	)
}
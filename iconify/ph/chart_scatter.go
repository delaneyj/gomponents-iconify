package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 208a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8V48a8 8 0 0 1 16 0v152h184a8 8 0 0 1 8 8Zm-100-48a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm-24-56a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm-32 72a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm96-48a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm24-40a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm-20 76a12 12 0 1 0 12-12a12 12 0 0 0-12 12Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatterLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 208a6 6 0 0 1-6 6H32a6 6 0 0 1-6-6V48a6 6 0 0 1 12 0v154h186a6 6 0 0 1 6 6Zm-98-50a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm-24-56a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm-32 72a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm96-48a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm24-40a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm-8 68a10 10 0 1 0 10 10a10 10 0 0 0-10-10Z"/>`),
		g.Group(children),
	)
}
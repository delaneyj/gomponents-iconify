package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatterBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 208a12 12 0 0 1-12 12H32a12 12 0 0 1-12-12V48a12 12 0 0 1 24 0v148h180a12 12 0 0 1 12 12Zm-120-60a16 16 0 1 0 16-16a16 16 0 0 0-16 16Zm-8-40a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm-32 72a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm96-48a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm24-40a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm-24 72a16 16 0 1 0 16-16a16 16 0 0 0-16 16Z"/>`),
		g.Group(children),
	)
}
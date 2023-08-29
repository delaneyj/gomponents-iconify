package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatterThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 208a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4V48a4 4 0 0 1 8 0v156h188a4 4 0 0 1 4 4Zm-96-52a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm-24-56a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm-32 72a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm96-48a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm24-40a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm-8 88a8 8 0 1 0-8-8a8 8 0 0 0 8 8Z"/>`),
		g.Group(children),
	)
}
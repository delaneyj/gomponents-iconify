package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.865 3l9.526 16.5a1 1 0 0 1-.866 1.5H2.473a1 1 0 0 1-.866-1.5L11.133 3a1 1 0 0 1 1.732 0Zm-8.66 16h15.588L12 5.5L4.205 19Zm6.794-3h2v2h-2v-2Zm0-7h2v5h-2V9Z"/>`),
		g.Group(children),
	)
}
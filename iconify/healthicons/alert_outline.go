package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M21 8a3 3 0 1 1 6 0v22a3 3 0 1 1-6 0V8Zm0 32a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}
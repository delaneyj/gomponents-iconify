package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10.5 7.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm0 0v1a2 2 0 1 0 4 0v-1a7 7 0 1 0-3 5.745"/>`),
		g.Group(children),
	)
}
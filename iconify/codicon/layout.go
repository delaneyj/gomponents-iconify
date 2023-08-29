package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 2L2 3v10l1 1h4l1-1V3L7 2H3Zm0 11V3h4v10H3Zm7-10l1-1h3l1 1v3l-1 1h-3l-1-1V3Zm1 0v3h3V3h-3Zm-1 7l1-1h3l1 1v3l-1 1h-3l-1-1v-3Zm1 0v3h3v-3h-3Z"/>`),
		g.Group(children),
	)
}
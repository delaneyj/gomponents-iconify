package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-1.07A7.001 7.001 0 0 1 5 13V8a7 7 0 0 1 14 0v5a7.001 7.001 0 0 1-6 6.93V21h4v2H7v-2h4Zm1-12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}
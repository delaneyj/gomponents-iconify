package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.617 4.076A1 1 0 0 0 12 5v3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h8v3a1 1 0 0 0 1.707.707l7-7a1 1 0 0 0 0-1.414l-7-7a1 1 0 0 0-1.09-.217zM18.586 12L14 16.586V15a1 1 0 0 0-1-1H5v-4h8a1 1 0 0 0 1-1V7.414L18.586 12z"/></g>`),
		g.Group(children),
	)
}
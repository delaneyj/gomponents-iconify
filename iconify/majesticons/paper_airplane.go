package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperAirplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2.245 21.655a1 1 0 0 1-.14-1.102l9-18a1 1 0 0 1 1.79 0l9 18a1 1 0 0 1-1.211 1.396L13 19.387V12a1 1 0 1 0-2 0v7.387L3.316 21.95a1 1 0 0 1-1.071-.294z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
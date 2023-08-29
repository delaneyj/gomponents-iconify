package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M15 4c1.306 0 2.418.835 2.83 2H9a5 5 0 0 0-5 5v4.83A3.001 3.001 0 0 1 2 13V7a3 3 0 0 1 3-3h10zm4 4H9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3zm-3 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
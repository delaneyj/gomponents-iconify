package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-4v2h1a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h1v-2H5a3 3 0 0 1-3-3V6zm11 11h-2v2h2v-2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 5a1 1 0 0 0-1 1v1h16V6a1 1 0 0 0-1-1H5zM2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1a2 2 0 0 1-1 1.732V18a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V8.732A2 2 0 0 1 2 7V6zm3 3v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V9H5zm4 3a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}
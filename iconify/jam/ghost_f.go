package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 0 1 8 8v12l-4.919-1l-3.08 1l-2.992-1L0 20V8a8 8 0 0 1 8-8zM5.5 10a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3z"/>`),
		g.Group(children),
	)
}
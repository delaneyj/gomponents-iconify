package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a1 1 0 0 0 0 2h1v13a5 5 0 0 0 10 0V4h1a1 1 0 1 0 0-2H6zm3 5V4h6v3H9zm4.708 5.708a1 1 0 0 1-1.414 0h-.001a1 1 0 0 1 1.414-1.415a1 1 0 0 1 0 1.415zm-3.414 4a1 1 0 1 0 1.414-1.414v-.001a1 1 0 0 0-1.415 1.414z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
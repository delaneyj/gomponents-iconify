package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v10h4V7H4zm6 0v10h4V7h-4zm6 0v10h4V7h-4z"/>`),
		g.Group(children),
	)
}
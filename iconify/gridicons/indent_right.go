package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4H4v16h2V4zm16 9H11.828l2.086 2.086L12.5 16.5L8 12l4.5-4.5l1.414 1.414L11.828 11H22v2z"/>`),
		g.Group(children),
	)
}
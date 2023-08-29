package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCheckLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.455 19L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455Zm-.692-2H20V5H4v13.385L5.763 17Zm5.53-4.879l4.243-4.242l1.414 1.414l-5.657 5.657l-3.89-3.89l1.415-1.414l2.475 2.475Z"/>`),
		g.Group(children),
	)
}
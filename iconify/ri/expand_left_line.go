package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.071 4.93l1.414 1.413L6.828 11H16v2H6.828l4.657 4.657l-1.414 1.414L3 12.001l7.071-7.072ZM18.001 19V5h2v14h-2Z"/>`),
		g.Group(children),
	)
}
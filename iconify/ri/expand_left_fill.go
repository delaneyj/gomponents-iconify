package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.414 4.586V11H16v2h-5.586v6.414L3 12l7.414-7.414ZM18 19V5h2v14h-2Z"/>`),
		g.Group(children),
	)
}
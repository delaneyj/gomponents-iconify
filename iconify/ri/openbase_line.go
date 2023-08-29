package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenbaseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L2 5.996l1 9.464l9 7.04l9-7.04l1-9.464L12 3Zm7.837 4.436L12 19.96L4.163 7.436L12 5.088l7.837 2.348Z"/>`),
		g.Group(children),
	)
}
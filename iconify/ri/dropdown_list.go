package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropdownList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5v14h9v2H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v5h-2V5H4Zm15.707 6.293L19 10.586l-.707.707l-3 3l1.414 1.414L19 13.414l2.293 2.293l1.414-1.414l-3-3Zm-3 6L19 19.586l2.293-2.293l1.414 1.414l-3 3l-.707.707l-.707-.707l-3-3l1.414-1.414Z"/>`),
		g.Group(children),
	)
}
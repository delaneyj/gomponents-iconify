package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsSquareDownF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm5.998 12.612l-2.122-2.121a1 1 0 1 0-1.414 1.414l2.829 2.828a1 1 0 0 0 1.414 0l2.828-2.828a1 1 0 0 0-1.414-1.414l-2.121 2.121zm0-5L7.876 5.491a1 1 0 1 0-1.414 1.414l2.829 2.828a1 1 0 0 0 1.414 0l2.828-2.828A1 1 0 0 0 12.12 5.49L9.998 7.612z"/>`),
		g.Group(children),
	)
}
package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.263 2.826h-6.36l-5.2 5.202v6.36L18.404 30.09l11.56-11.562l-15.7-15.702zM6.495 8.86a1.581 1.581 0 0 1 0-2.24c.62-.62 1.622-.62 2.24 0c.62.62.62 1.62 0 2.24c-.618.62-1.62.62-2.24 0z"/>`),
		g.Group(children),
	)
}
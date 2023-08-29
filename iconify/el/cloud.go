package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M983.888 575.377c187.925-18.507 293.084 231.644 148.656 358.546H49.759C-89.529 697.252 82.314 382.276 333.563 443.401C535.007 115.536 908.131 291.199 983.88 575.377h.008z"/>`),
		g.Group(children),
	)
}
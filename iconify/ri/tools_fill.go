package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.329 3.272A3.5 3.5 0 0 1 9.8 7.745L20.646 18.59l-2.121 2.122L7.679 9.867a3.5 3.5 0 0 1-4.472-4.474L5.443 7.63a1.5 1.5 0 0 0 2.122-2.121L5.329 3.272Zm10.367 1.883l3.182-1.768l1.415 1.415l-1.768 3.182l-1.768.353l-2.121 2.121l-1.415-1.414l2.122-2.121l.353-1.768Zm-7.07 7.778l2.12 2.122l-4.95 4.95a1.5 1.5 0 0 1-2.218-2.015l.097-.107l4.95-4.95Z"/>`),
		g.Group(children),
	)
}
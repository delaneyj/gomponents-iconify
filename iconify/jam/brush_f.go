package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.36 11.116l2.68 2.681c-.5.65-1.122 1.26-1.864 1.832c-1.492 1.15-3.265 1.445-5.318.885a82.957 82.957 0 0 1 2.603-3.6c.489-.637 1.121-1.236 1.898-1.798zm1.489-1.338c.55-.79 1.269-1.591 2.154-2.406c1.266-1.165 3.467-3.124 6.602-5.877c.834-.733 1.86-.9 2.375-.385c.514.514.348 1.542-.384 2.376c-2.757 3.14-4.716 5.34-5.878 6.601c-.814.883-1.616 1.6-2.407 2.153L6.849 9.778z"/>`),
		g.Group(children),
	)
}
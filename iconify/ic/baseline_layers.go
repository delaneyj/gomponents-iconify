package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineLayers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.99 18.54l-7.37-5.73L3 14.07l9 7l9-7l-1.63-1.27l-7.38 5.74zM12 16l7.36-5.73L21 9l-9-7l-9 7l1.63 1.27L12 16z"/>`),
		g.Group(children),
	)
}
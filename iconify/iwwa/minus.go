package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="m36.495 19.226l-32.99.024a.763.763 0 0 0 0 1.525l32.99-.023a.764.764 0 0 0 0-1.526z"/>`),
		g.Group(children),
	)
}
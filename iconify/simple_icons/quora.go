package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.38.948A11.963 11.963 0 0 1 21.248 19.54l2.41 2.422c.732.736.21 1.99-.828 1.99l-10.71.01a12.52 12.52 0 0 1-.304 0h-.02A11.963 11.963 0 0 1 7.382.95Zm7.322 4.428a7.172 7.172 0 1 0-5.488 13.252a7.172 7.172 0 0 0 5.489-13.252Z"/>`),
		g.Group(children),
	)
}
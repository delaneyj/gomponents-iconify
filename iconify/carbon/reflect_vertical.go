package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16.79 19.386l7 9A1 1 0 0 1 23 30H9a1 1 0 0 1-.79-1.614l7-9a1 1 0 0 1 1.58 0ZM2 17v-2h28v2zm14-4a1.001 1.001 0 0 1-.79-.386l-7-9A1 1 0 0 1 9 2h14a1 1 0 0 1 .79 1.614l-7 9A1.001 1.001 0 0 1 16 13Zm-4.956-9L16 10.371L20.956 4Z"/>`),
		g.Group(children),
	)
}
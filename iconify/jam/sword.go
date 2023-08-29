package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.778 10.728l1.414 1.414l8.436-8.435l-.236-1.179l-1.178-.236l-8.436 8.436zM3.95 16.385a1 1 0 0 1-1.414 1.414L1.12 16.385a1 1 0 0 1 1.415-1.414l2.828-2.829l-1.414-1.414a1 1 0 0 1 1.414-1.414L14.556.12l3.536.707l.707 3.536l-9.192 9.192a1 1 0 1 1-1.415 1.415l-1.414-1.415l-2.828 2.829z"/>`),
		g.Group(children),
	)
}
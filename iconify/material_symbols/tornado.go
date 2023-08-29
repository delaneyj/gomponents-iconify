package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22l-2.9 5H3.9L1 3Zm4.05 7h13.9l-1.75 3H6.8l-1.75-3Zm2.9 5h8.1L12 22l-4.05-7Z"/>`),
		g.Group(children),
	)
}
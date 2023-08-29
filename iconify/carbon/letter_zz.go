package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterZz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 15v-2h-8v2h5.5L17 21v2h8v-2h-5.49L25 15zM15 9H7v2h6L7 21v2h8v-2H9l6-10V9z"/>`),
		g.Group(children),
	)
}
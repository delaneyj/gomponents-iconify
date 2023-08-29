package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M297 141q6-7 5-16t-7-14L152 0L11 111q-15 15-2 30q16 14 30 2l92-77v297q0 21 21 21t21-21V66l94 77q17 15 30-2z"/>`),
		g.Group(children),
	)
}
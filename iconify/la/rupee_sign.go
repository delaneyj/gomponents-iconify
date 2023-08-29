package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RupeeSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 5v2h4c1.703 0 3.941 1.04 4.719 3H8v2h8.969c-.309 2.61-2.996 4-4.969 4H8v2.469L18.25 27h3.125l-10.813-9H12c3.234 0 6.676-2.39 6.969-6H24v-2h-5.188C18.509 8.816 17.86 7.805 17 7h7V5z"/>`),
		g.Group(children),
	)
}
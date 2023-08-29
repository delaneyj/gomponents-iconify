package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rupee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 5v2h4c1.704 0 3.94 1.038 4.72 3H8v2h8.97c-.31 2.61-2.996 4-4.97 4H8v2.47L18.25 27h3.125l-10.813-9H12c3.234 0 6.674-2.39 6.97-6H24v-2h-5.188C18.51 8.816 17.86 7.804 17 7h7V5H8z"/>`),
		g.Group(children),
	)
}
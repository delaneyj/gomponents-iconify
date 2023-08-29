package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.336 9.685a9.934 9.934 0 0 0 11.256 12.123A9.931 9.931 0 0 0 20.708 7.23A10.046 10.046 0 0 0 12 2.072a.507.507 0 0 0-.5.5v4.2a.5.5 0 0 0 1 0v-4.2l-.5.5a8.935 8.935 0 0 1 8.433 11.892a8.938 8.938 0 0 1-13.965 4.063A9.041 9.041 0 0 1 3.3 9.951c.142-.627-.822-.9-.964-.266Z"/><path fill="currentColor" d="M7.4 8.117a.5.5 0 0 1 .707-.707l4.243 4.242a.5.5 0 0 1-.707.707Z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 4v9h16V4H0zm10 2h1v1h-1V6zM8 6h1v1H8V6zm2 2v1H9V8h1zM6 6h1v1H6V6zm2 2v1H7V8h1zM4 6h1v1H4V6zm2 2v1H5V8h1zM2 6h1v1H2V6zm1 5H2v-1h1v1zm0-3h1v1H3V8zm9 3H4v-1h8v1zm0-2h-1V8h1v1zm2 2h-1v-1h1v1zm0-2h-1V7h-1V6h2v3z"/>`),
		g.Group(children),
	)
}
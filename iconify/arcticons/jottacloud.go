package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jottacloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.056 5.5H42.5v24a13 13 0 0 1-13 13h-3.444h0v-24a13 13 0 0 1 13-13ZM5.5 26.056h3.444a13 13 0 0 1 13 13V42.5h0H18.5a13 13 0 0 1-13-13v-3.444h0Z"/>`),
		g.Group(children),
	)
}
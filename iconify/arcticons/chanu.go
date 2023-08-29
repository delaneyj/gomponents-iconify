package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chanu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.75 24A9.25 9.25 0 1 1 24 14.75V24ZM24 14.75A9.25 9.25 0 1 1 33.25 24H24ZM33.25 24A9.25 9.25 0 1 1 24 33.25V24ZM24 33.25A9.25 9.25 0 1 1 14.75 24H24Z"/>`),
		g.Group(children),
	)
}
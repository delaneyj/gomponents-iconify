package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Directions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5h11l2 2l-2 2h-11a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1zm12 7h-11l-2 2l2 2h11a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1zm-6-3v3m0 4v3"/>`),
		g.Group(children),
	)
}
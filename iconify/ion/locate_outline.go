package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocateOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="48" d="M256 96V56m0 400v-40"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="M256 112a144 144 0 1 0 144 144a144 144 0 0 0-144-144Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="48" d="M416 256h40m-400 0h40"/>`),
		g.Group(children),
	)
}
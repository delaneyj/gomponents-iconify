package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocateSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="48" d="M256 96V56m0 400v-40m0-304a144 144 0 1 0 144 144a144 144 0 0 0-144-144Zm160 144h40m-400 0h40"/>`),
		g.Group(children),
	)
}
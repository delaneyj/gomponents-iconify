package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistMusicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6v2H3V6h12m0 4v2H3v-2h12M3 16v-2h8v2H3M17 6h5v2h-3v9a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3c.35 0 .69.07 1 .18V6m-1 10a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
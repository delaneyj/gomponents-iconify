package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueueMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20q-1.25 0-2.125-.875T13 17q0-1.25.875-2.125T16 14q.275 0 .525.038T17 14.2V6h5v2h-3v9q0 1.25-.875 2.125T16 20ZM3 16v-2h8v2H3Zm0-4v-2h12v2H3Zm0-4V6h12v2H3Z"/>`),
		g.Group(children),
	)
}
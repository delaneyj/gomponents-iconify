package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 192h256v32H128zm0 112h128v32H128z"/><path fill="currentColor" d="M48 432h416V88H48Zm32-312h352v280H80Z"/>`),
		g.Group(children),
	)
}
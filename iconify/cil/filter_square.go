package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 16H40a24.027 24.027 0 0 0-24 24v432a24.027 24.027 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24Zm-8 448H48V48h416Z"/><path fill="currentColor" d="M208 427h48.627L301 368.076V260.87l114-129.018V85H89v47.176l119 129.018Zm-87-307.328V117h262v2.739L269 248.757v108.618L240.666 395H240V248.69Z"/>`),
		g.Group(children),
	)
}
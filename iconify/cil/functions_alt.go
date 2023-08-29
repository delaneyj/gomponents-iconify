package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 416h232v-32H160v-9.717L269.834 256L160 137.717V128h200V96H128v54.283L226.166 256L128 361.717V416z"/>`),
		g.Group(children),
	)
}
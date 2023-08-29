package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 16v480h32V360h352v-37.238L363.841 208L448 93.238V56H96V16Zm348.159 72l-88 120l88 120H96V88Z"/>`),
		g.Group(children),
	)
}
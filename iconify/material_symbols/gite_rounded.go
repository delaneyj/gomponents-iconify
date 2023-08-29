package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.825 0-1.413-.588T2 17v-6.175q0-.4.15-.762t.425-.638l2.85-2.85q.275-.275.638-.425T6.825 6H7V5q0-.425.288-.712T8 4q.425 0 .713.288T9 5v1h8.175q.4 0 .763.15t.637.425l2.85 2.85q.275.275.425.638t.15.762V17q0 .825-.588 1.413T20 19H4Zm12-2h4v-6.175l-2-2l-2 2V17ZM4 17h10v-5H4v5Z"/>`),
		g.Group(children),
	)
}
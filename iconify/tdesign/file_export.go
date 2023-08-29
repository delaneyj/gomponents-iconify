package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h6v2H3V1Zm12 2.414V7h3.586L15 3.414Zm4.05 10.674l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508h-7.11v-2h7.11l-2.48-2.508l1.422-1.406Z"/>`),
		g.Group(children),
	)
}
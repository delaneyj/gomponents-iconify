package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageDocumentSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293v6.999l-2.5-2.504l-2.959 2.957L4.5 5.7L1 10.075V1.5ZM11 6h-1V5h1v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M1 11.675V13.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-1.793l-2.5-2.504l-3.041 3.039L4.5 7.3L1 11.675Z"/>`),
		g.Group(children),
	)
}
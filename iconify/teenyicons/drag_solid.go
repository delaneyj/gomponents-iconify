package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1H0V0h1v1Zm4 0H4V0h1v1Zm4 0H8V0h1v1ZM3 3h11v4h-1V4H4v9h3v1H3V3ZM1 5H0V4h1v1Zm6.146 2.146a.5.5 0 0 1 .491-.127l7 2a.5.5 0 0 1 .06.94l-3.316 1.422l-1.421 3.316a.5.5 0 0 1-.94-.06l-2-7a.5.5 0 0 1 .126-.49Zm1.082 1.082l1.366 4.782l.946-2.207a.5.5 0 0 1 .263-.263l2.207-.946l-4.782-1.366ZM1 9H0V8h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
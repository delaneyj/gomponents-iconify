package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MpFourSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 8h.5a.5.5 0 0 0 0-1H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM6 6h1.5a1.5 1.5 0 1 1 0 3H7v2H6V6Zm-3.691.038a.5.5 0 0 1 .545.108l.646.647l.646-.647A.5.5 0 0 1 5 6.5V11H4V7.707l-.5.5l-.5-.5V11H2V6.5a.5.5 0 0 1 .309-.462ZM11 6h-1v3h2v2h1V6h-1v2h-1V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
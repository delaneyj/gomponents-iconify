package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v4H0V1.5ZM0 5v8.5A1.5 1.5 0 0 0 1.5 15H4V5H0Zm5 10h8.5a1.5 1.5 0 0 0 1.5-1.5V5H5v10ZM15 4V1.5A1.5 1.5 0 0 0 13.5 0H5v4h10Z"/>`),
		g.Group(children),
	)
}
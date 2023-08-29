package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m4.5 4.5l6 6m-6 0l6-6m-3 10a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		g.Group(children),
	)
}
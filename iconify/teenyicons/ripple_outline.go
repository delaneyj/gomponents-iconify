package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RippleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m1.5 1.5l3.06 3.316a4 4 0 0 0 5.88 0L13.5 1.5m-12 12l3.06-3.316a4 4 0 0 1 5.88 0L13.5 13.5"/>`),
		g.Group(children),
	)
}
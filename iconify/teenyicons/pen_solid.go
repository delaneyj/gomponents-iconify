package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2 0h11v.5a5.482 5.482 0 0 1-1.874 4.134l1.968 3.93L8 14.672V8.994H7v5.678L1.907 8.564l1.967-3.93A5.482 5.482 0 0 1 2 .5V0Z"/>`),
		g.Group(children),
	)
}
package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="m560 1l200 200l-1 800H0V1h560zM160 841h440V281H480V161H160v680z"/>`),
		g.Group(children),
	)
}
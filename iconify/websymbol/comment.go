package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1137 101v532q0 41-29.5 70.5T1037 733H869v268L602 733H100q-41 0-70.5-29.5T0 633V101q0-41 29.5-70.5T100 1h937q41 0 70.5 29.5T1137 101z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1 10.5h13m-11.5 0v-5a5 5 0 0 1 10 0v5m-7 1.5v.5a2 2 0 1 0 4 0V12"/>`),
		g.Group(children),
	)
}
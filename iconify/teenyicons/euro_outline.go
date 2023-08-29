package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13.374 3A6 6 0 0 0 2.5 6.5v2A6 6 0 0 0 13.374 12M0 5.5h7m-7 4h7"/>`),
		g.Group(children),
	)
}
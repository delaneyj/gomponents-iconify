package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 1.5h1.5a6 6 0 1 1 0 12H0m7-12h4.5a3 3 0 1 1 0 6m0 0H9m2.5 0h-2m2 0a3 3 0 1 1 0 6H7"/>`),
		g.Group(children),
	)
}
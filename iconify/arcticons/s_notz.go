package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SNotz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h25.56V32.56H43V7a2 2 0 0 0-2-2Zm-8.44 38L43 32.56M10.5 14h27m-27 10h27m-27 10h16.06"/>`),
		g.Group(children),
	)
}
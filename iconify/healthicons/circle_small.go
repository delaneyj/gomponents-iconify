package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M24 36c6.627 0 12-5.373 12-12c0-6.628-5.373-12-12-12s-12 5.372-12 12c0 6.627 5.373 12 12 12Z"/>`),
		g.Group(children),
	)
}
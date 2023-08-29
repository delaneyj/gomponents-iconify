package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func F(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 12a2 2 0 0 1 2-2h12a2 2 0 1 1 0 4H20v8h10a2 2 0 1 1 0 4H20v10a2 2 0 1 1-4 0V12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
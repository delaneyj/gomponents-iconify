package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Off(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M459 83c118 50 201 167 201 304c0 181-148 330-330 330C149 717 0 568 0 387C0 250 83 133 202 83v111C140 236 98 306 98 387c0 128 104 232 232 232s232-104 232-232c0-81-41-151-103-193V83zm-74 302V0H275v385h110z"/>`),
		g.Group(children),
	)
}
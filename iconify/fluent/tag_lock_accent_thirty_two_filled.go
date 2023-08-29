package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagLockAccentThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.684 4a3 3 0 0 0-2.05.81L4.455 16.203a3.25 3.25 0 0 0-.078 4.672l6.326 6.326a3.25 3.25 0 0 0 4.298.264V22a4.002 4.002 0 0 1 3.08-3.894a5.002 5.002 0 0 1 8.697-2.383l.344-.344A3 3 0 0 0 28 13.257V6.5A2.5 2.5 0 0 0 25.5 4h-6.816Z"/>`),
		g.Group(children),
	)
}
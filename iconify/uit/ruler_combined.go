package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCombined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V10h11.5a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5zM21 9h-3.5V6.5a.5.5 0 0 0-1 0V9h-3V6.5a.5.5 0 0 0-1 0V9H10V6.5a.5.5 0 0 0-1 0V9H6.5a.5.5 0 0 0 0 1H9v2.5H6.5a.5.5 0 0 0 0 1H9v3H6.5a.5.5 0 0 0 0 1H9V21H3V3h18v6z"/>`),
		g.Group(children),
	)
}
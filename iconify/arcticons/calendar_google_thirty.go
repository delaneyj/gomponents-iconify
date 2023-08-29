package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.692 25.5a2.987 2.987 0 0 0 3 3a2.895 2.895 0 0 0 2.89-3v-3a2.966 2.966 0 0 0-2.89-3a3.06 3.06 0 0 0-3 3Zm-3.607-1.556a2.229 2.229 0 0 1 2.223 2.223h0a2.229 2.229 0 0 1-2.223 2.222h-.889a3.614 3.614 0 0 1-2.777-.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.419 20.278a3.78 3.78 0 0 1 2.778-.778h.888a2.229 2.229 0 0 1 2.223 2.222h0a2.229 2.229 0 0 1-2.223 2.222h-2.222"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diceware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="34" cy="12" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.596"/><circle cx="34" cy="24" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.596"/><circle cx="34" cy="36" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.596"/><circle cx="14" cy="12" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.596"/><circle cx="14" cy="24" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.596"/><circle cx="14" cy="36" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.596"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Agenda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31" height="8.5" x="8.5" y="4.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="31" height="8.5" x="8.5" y="17.125" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="31" height="14" x="8.5" y="29.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><circle cx="34.75" cy="9" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.75" cy="21.375" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.75" cy="33.75" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
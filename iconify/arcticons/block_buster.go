package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockBuster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="29.514" height="29.514" x="9.38" y="9.243" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.216" ry="4.216"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.662 10.477l-27.05 27.048M4.5 24.142h39"/>`),
		g.Group(children),
	)
}
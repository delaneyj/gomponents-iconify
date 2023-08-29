package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quickswitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="17.249" height="33.211" x="15.375" y="7.395" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.065"/><rect width="7.195" height="28.398" x="4.5" y="9.8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.065"/><rect width="7.195" height="28.398" x="36.305" y="9.8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.065"/>`),
		g.Group(children),
	)
}
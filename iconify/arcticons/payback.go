package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Payback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="15.787" cy="16.637" r="6.975" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.438" cy="16.637" r="6.975" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.787" cy="32.837" r="6.975" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.438" cy="32.837" r="6.975" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
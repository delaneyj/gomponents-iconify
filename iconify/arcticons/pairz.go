package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pairz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="17.61" height="17.61" x="4.5" y="15.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.21"/><rect width="17.61" height="17.61" x="25.89" y="15.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.21"/>`),
		g.Group(children),
	)
}
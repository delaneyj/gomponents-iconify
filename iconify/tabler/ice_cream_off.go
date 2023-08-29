package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21.5V17M8 8v9h8v-1m0-4V7a4 4 0 0 0-7.277-2.294M8 10.5l1.74-.76m2.79-1.222L16 7m-8 7.5l4.488-1.964M3 3l18 18"/>`),
		g.Group(children),
	)
}
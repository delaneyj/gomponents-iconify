package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21.5V17m-4 0h8V7a4 4 0 1 0-8 0v10zm0-6.5L16 7m-8 7.5l8-3.5"/>`),
		g.Group(children),
	)
}
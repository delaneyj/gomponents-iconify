package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18m-5.788-5.79L11 21v-7H5l3.79-5.21m1.685-2.32L13 3v6m1 1h5l-2.104 2.893"/>`),
		g.Group(children),
	)
}
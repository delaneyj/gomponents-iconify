package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewfinderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.65 5.623a9 9 0 1 0 12.71 12.745m1.684-2.328A9 9 0 0 0 7.95 3.96M12 3v4m0 14v-3m-9-6h4m14 0h-3m-6 0v.01M3 3l18 18"/>`),
		g.Group(children),
	)
}
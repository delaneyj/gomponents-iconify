package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShadowOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.634 5.638a9 9 0 0 0 12.728 12.727m1.68-2.32A9 9 0 0 0 7.956 3.957M16 12h2m-5 3h2m-2 3h1m-1-9h4m-4-3h1M3 3l18 18"/>`),
		g.Group(children),
	)
}
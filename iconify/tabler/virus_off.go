package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18M8.469 8.46a5 5 0 0 0 7.058 7.084m1.386-2.608a5 5 0 0 0-5.826-5.853M12 7V3m-1 0h2m2.536 5.464l2.828-2.828m-.707-.707l1.414 1.414M17 12h4m0-1v2m-2.636 5.363l-.707.707M12 17v4m1 0h-2m-2.535-5.464l-2.829 2.828m.707.707L4.93 17.657M7 12H3m0 1v-2m2.636-5.363l-.707.707"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioactiveOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.118 14.127c-.182.181-.39.341-.618.473l3 5.19a9 9 0 0 0 1.856-1.423m1.68-2.32A8.993 8.993 0 0 0 21 12h-5m-2.5-2.6l3-5.19a9 9 0 0 0-8.536-.25M10.5 14.6l-3 5.19A9 9 0 0 1 3 12h6a3 3 0 0 0 1.5 2.6M3 3l18 18"/>`),
		g.Group(children),
	)
}
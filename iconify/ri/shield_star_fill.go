package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldStarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.783 2.826L12 1l8.217 1.826a1 1 0 0 1 .783.976v9.987a6 6 0 0 1-2.672 4.992L12 23l-6.328-4.219A6 6 0 0 1 3 13.79V3.802a1 1 0 0 1 .783-.976ZM12 13.5l2.939 1.545l-.561-3.272l2.377-2.318l-3.285-.478L12 6l-1.47 2.977l-3.285.478l2.377 2.318l-.56 3.272L12 13.5Z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5.5l4.286 4.142a5.737 5.737 0 0 1 1.607 2.971a5.62 5.62 0 0 1-.362 3.334a5.85 5.85 0 0 1-2.21 2.584A6.15 6.15 0 0 1 7.5 14.5a6.15 6.15 0 0 1-3.32-.97a5.85 5.85 0 0 1-2.211-2.583a5.62 5.62 0 0 1-.363-3.334a5.737 5.737 0 0 1 1.608-2.97L7.5.5Z"/>`),
		g.Group(children),
	)
}
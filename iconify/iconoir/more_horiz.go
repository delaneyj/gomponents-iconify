package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHoriz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-8 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-8 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}
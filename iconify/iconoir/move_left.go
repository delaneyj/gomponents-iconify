package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-5-2H2m0 0l3-3m-3 3l3 3"/>`),
		g.Group(children),
	)
}
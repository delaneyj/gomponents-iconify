package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyCommand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6v12m6-12v12M9 6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3M9 18a3 3 0 1 1-3-3h12a3 3 0 1 1-3 3"/>`),
		g.Group(children),
	)
}
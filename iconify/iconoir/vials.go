package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vials(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21H3m6-9H5m14 0h-4m-8 6a2 2 0 0 1-2-2V3h4v13a2 2 0 0 1-2 2Zm10 0a2 2 0 0 1-2-2V3h4v13a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}
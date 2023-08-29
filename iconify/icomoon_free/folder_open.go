package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m13 15l3-8H3l-3 8zM2 6l-2 9V2h4.5l2 2H13v2z"/>`),
		g.Group(children),
	)
}
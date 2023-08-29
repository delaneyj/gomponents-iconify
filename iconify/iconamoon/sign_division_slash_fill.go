package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDivisionSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.447 3.106a1 1 0 0 1 .447 1.341l-8 16a1 1 0 1 1-1.788-.894l8-16a1 1 0 0 1 1.341-.447Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
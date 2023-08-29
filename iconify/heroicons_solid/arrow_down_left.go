package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.78 3.22a.75.75 0 0 0-1.06 0L4.5 14.44V6.75a.75.75 0 0 0-1.5 0v9.5c0 .414.336.75.75.75h9.5a.75.75 0 0 0 0-1.5H5.56L16.78 4.28a.75.75 0 0 0 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
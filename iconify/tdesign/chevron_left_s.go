package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.315 12l4.95-4.95l1.414 1.415L11.144 12l3.535 3.536l-1.414 1.414L8.315 12Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17.909l-4.95-4.95l1.415-1.414l3.536 3.535l3.535-3.535l1.414 1.414l-4.95 4.95Zm0-5.461l-4.95-4.95l1.415-1.414l3.536 3.535l3.535-3.535l1.414 1.414l-4.95 4.95Z"/>`),
		g.Group(children),
	)
}
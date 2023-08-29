package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m.586 12l4.95-4.95L6.95 8.464L3.414 12l3.536 3.535l-1.414 1.414L.586 12Zm16.464 3.535L20.586 12L17.05 8.464l1.415-1.414l4.95 4.95l-4.95 4.95l-1.415-1.415Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.684 12l-4.95 4.95l-1.414-1.414L12.856 12L9.32 8.465l1.415-1.415l4.95 4.95Z"/>`),
		g.Group(children),
	)
}
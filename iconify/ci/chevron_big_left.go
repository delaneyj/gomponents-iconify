package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronBigLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.535 3.515L7.05 12l8.485 8.485l1.415-1.414L9.878 12l7.072-7.071l-1.415-1.414Z"/>`),
		g.Group(children),
	)
}
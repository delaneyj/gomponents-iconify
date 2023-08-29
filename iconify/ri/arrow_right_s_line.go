package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightSLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.171 12l-4.95-4.95l1.415-1.413L16 12l-6.364 6.364l-1.414-1.415l4.95-4.95Z"/>`),
		g.Group(children),
	)
}
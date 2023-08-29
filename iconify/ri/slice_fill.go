package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.768 12.229l2.121 2.121c-4.596 4.596-10.253 6.01-13.788 5.304L17.657 4.097l2.121 2.121l-6.01 6.01Z"/>`),
		g.Group(children),
	)
}
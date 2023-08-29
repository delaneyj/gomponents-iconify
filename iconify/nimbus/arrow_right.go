package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.85 6.69L7.14 1.55l-.7 1l7.12 4.75H.5v1.28h12.76l-6.83 4.86l.72 1L14.85 9a1.42 1.42 0 0 0 .65-1.15a1.4 1.4 0 0 0-.65-1.16z"/>`),
		g.Group(children),
	)
}
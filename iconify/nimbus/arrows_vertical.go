package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.12 6.75H4.87v7.05l-2.05-3l-1 .71l2.67 3.93a1.29 1.29 0 0 0 1 .59a1.29 1.29 0 0 0 1-.59l2.67-3.93l-1-.71l-2.06 3zM9.45.59L6.78 4.52l1 .71l2.06-3v7.02h1.25v-7l2 3l1-.71L11.55.59a1.23 1.23 0 0 0-2.1 0z"/>`),
		g.Group(children),
	)
}
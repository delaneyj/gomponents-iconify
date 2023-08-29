package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareForwardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14h-2a8.999 8.999 0 0 0-7.968 4.81A10.133 10.133 0 0 1 3 18C3 12.477 7.477 8 13 8V3l10 8l-10 8v-5Z"/>`),
		g.Group(children),
	)
}
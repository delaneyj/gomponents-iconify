package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CelsiusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 10a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM22 10h-2a4 4 0 0 0-8 0v5a4 4 0 0 0 8 0h2a6 6 0 0 1-12 0v-5a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.374 19.44a1.5 1.5 0 1 1-2.121 2.12l-4.596-4.595L12.12 20.5l-7.778-7.778A8 8 0 0 1 4.17 1.587l.173-.179L22.374 19.44Z"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 10.27l-5-2.89a2 2 0 0 0-3 1.73v5.78a2 2 0 0 0 1 1.73a2 2 0 0 0 2 0l5-2.89a2 2 0 0 0 0-3.46ZM15 12l-5 2.89V9.11L15 12ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}
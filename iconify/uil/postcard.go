package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 11h1a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1ZM6 12h5a1 1 0 0 0 0-2H6a1 1 0 0 0 0 2Zm16-8H2a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Zm-1 14H3V6h18ZM6 16h5a1 1 0 0 0 0-2H6a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}
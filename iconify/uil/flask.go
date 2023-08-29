package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.11 17.49L15 8.73V4h1a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2h1v4.73l-5.11 8.76A3 3 0 0 0 6.48 22h11a3 3 0 0 0 2.59-4.51Zm-9.25-8A1 1 0 0 0 11 9V4h2v5a1 1 0 0 0 .14.5L14 11h-4Zm7.52 10a1 1 0 0 1-.86.5h-11a1 1 0 0 1-.86-.5a1 1 0 0 1 0-1L8.83 13h6.35l3.2 5.5a1 1 0 0 1 0 1ZM10 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 1a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
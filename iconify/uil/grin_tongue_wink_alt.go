package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinTongueWinkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.21 10.54a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 1 0 1.41 1.41a1 1 0 0 1 1.42 0ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm3-11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4H9a1 1 0 0 0 0 2a3 3 0 0 0 6 0a1 1 0 0 0 0-2Zm-3 3a1 1 0 0 1-1-1h2a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}
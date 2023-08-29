package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Surprise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8ZM10 9a1 1 0 1 0-1 1a1 1 0 0 0 1-1Zm5-1a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3 3a3.39 3.39 0 0 0-3.25 3.5A3.39 3.39 0 0 0 12 18a3.39 3.39 0 0 0 3.25-3.5A3.39 3.39 0 0 0 12 11Zm0 5a1.39 1.39 0 0 1-1.25-1.5A1.39 1.39 0 0 1 12 13a1.39 1.39 0 0 1 1.25 1.5A1.39 1.39 0 0 1 12 16Z"/>`),
		g.Group(children),
	)
}
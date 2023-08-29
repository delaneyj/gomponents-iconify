package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm12-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-16a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8 0a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1Zm-4 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
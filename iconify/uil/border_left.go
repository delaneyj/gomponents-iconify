package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-8-8a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1Zm16 2a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-8 2a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-16a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm12 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-8-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
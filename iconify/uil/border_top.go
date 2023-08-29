package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-8-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-14h16a1 1 0 0 0 0-2H4a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-8-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
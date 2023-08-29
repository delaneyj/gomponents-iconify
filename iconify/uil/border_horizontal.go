package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1ZM4 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm8 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1ZM4 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm12 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-4 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm8 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1ZM4 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm16 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4H4a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2Zm-4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM4 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm16-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM8 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
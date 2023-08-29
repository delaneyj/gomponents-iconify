package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1ZM4 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM20 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-4-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM4 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm16 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm5-7a1 1 0 0 0-1-1h-7V4a1 1 0 0 0-2 0v7H4a1 1 0 0 0 0 2h7v7a1 1 0 0 0 2 0v-7h7a1 1 0 0 0 1-1ZM8 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
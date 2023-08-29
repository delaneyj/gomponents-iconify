package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsByLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm1 4a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v4h1.5v4h3v-4H15v-4Zm-3-9C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2ZM4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0Z"/>`),
		g.Group(children),
	)
}
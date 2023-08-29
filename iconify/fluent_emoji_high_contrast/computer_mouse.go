package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerMouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 7a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v18a6 6 0 0 1-6 6h-8a6 6 0 0 1-6-6V7Zm6-4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V7a4 4 0 0 0-4-4h-3v3a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1V3h-3Z"/>`),
		g.Group(children),
	)
}
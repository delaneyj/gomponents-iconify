package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridVerticalRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4v-4h4v4Zm-8 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm8-6h-4V4h4v4Zm-6 0H7V4h4v4Z"/>`),
		g.Group(children),
	)
}
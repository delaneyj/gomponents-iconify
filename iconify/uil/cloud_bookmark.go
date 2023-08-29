package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudBookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11h-5a1 1 0 0 0-1 1v8a1 1 0 0 0 1.56.83l1.94-1.3l1.89 1.26A1 1 0 0 0 15 21a1 1 0 0 0 .44-.1A1 1 0 0 0 16 20v-8a1 1 0 0 0-1-1Zm-1 7.12l-.94-.63a1 1 0 0 0-1.12 0l-.94.64V13h3Zm4.42-10.9A7 7 0 0 0 5.06 9.11A4 4 0 0 0 6 17a1 1 0 0 0 0-2a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12a5 5 0 0 0-3.58-4.78Z"/>`),
		g.Group(children),
	)
}
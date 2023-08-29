package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11a1 1 0 0 1-1-1V7a5.002 5.002 0 0 1 8.532-3.542a5.09 5.09 0 0 1 1.306 2.293a1 1 0 0 1-1.934.505l-.002-.007a3.072 3.072 0 0 0-.786-1.379A3.002 3.002 0 0 0 9 7v3a1 1 0 0 1-1 1zm4 7a1 1 0 0 1-1-1v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M17 9H7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm-4 8a1 1 0 0 1-2 0v-3a1 1 0 1 1 2 0v3z"/>`),
		g.Group(children),
	)
}
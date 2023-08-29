package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmotionUnhappyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-5-3a5 5 0 0 1 10 0h-2a3 3 0 1 0-6 0H7Zm1-6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm8 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}
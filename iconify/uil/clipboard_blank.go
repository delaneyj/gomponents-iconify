package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardBlank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4h-1a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Zm-7 0h4v2h-4V4Zm8 15a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2h1a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}
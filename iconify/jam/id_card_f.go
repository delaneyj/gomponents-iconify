package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 0h16c1.105 0 2 .831 2 1.857v9.286C20 12.169 19.105 13 18 13H2c-1.105 0-2-.831-2-1.857V1.857C0 .831.895 0 2 0zm9 3a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-6zm0 3a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-6zM3 3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H3z"/>`),
		g.Group(children),
	)
}
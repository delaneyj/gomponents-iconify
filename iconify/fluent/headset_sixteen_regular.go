package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 6a4 4 0 1 1 8 0h-2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1a2 2 0 0 0 2-2V6A5 5 0 0 0 3 6v5a3 3 0 0 0 3 3h.585a1.5 1.5 0 1 0 0-1H6a2 2 0 0 1-2-2h2a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H4Zm8 1v2a1 1 0 0 1-1 1h-1V7h2Zm-6 3H4V7h2v3Zm1.5 3.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/>`),
		g.Group(children),
	)
}
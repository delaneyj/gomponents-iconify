package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0ZM1.5 8a6.5 6.5 0 1 0 13 0a6.5 6.5 0 0 0-13 0Zm9.78-2.22l-5.5 5.5a.749.749 0 0 1-1.275-.326a.749.749 0 0 1 .215-.734l5.5-5.5a.751.751 0 0 1 1.042.018a.751.751 0 0 1 .018 1.042Z"/>`),
		g.Group(children),
	)
}
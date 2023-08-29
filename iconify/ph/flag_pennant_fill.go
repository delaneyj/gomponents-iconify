package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPennantFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 104a8 8 0 0 1-5.37 7.56L64 173.69V216a8 8 0 0 1-16 0V40a8 8 0 0 1 10.63-7.56l184 64A8 8 0 0 1 248 104Z"/>`),
		g.Group(children),
	)
}
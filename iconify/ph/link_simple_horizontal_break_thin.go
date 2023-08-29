package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSimpleHorizontalBreakThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M20 128a44.05 44.05 0 0 0 44 44h40a4 4 0 0 1 0 8H64a52 52 0 0 1 0-104h40a4 4 0 0 1 0 8H64a44.05 44.05 0 0 0-44 44Zm172-52h-40a4 4 0 0 0 0 8h40a44 44 0 0 1 0 88h-40a4 4 0 0 0 0 8h40a52 52 0 0 0 0-104Z"/>`),
		g.Group(children),
	)
}
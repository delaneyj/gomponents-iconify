package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 232a8 8 0 0 1-8 8H56a8 8 0 0 1 0-16h144a8 8 0 0 1 8 8ZM80 192V40a16 16 0 0 1 16-16h64a16 16 0 0 1 16 16v152a16 16 0 0 1-16 16H96a16 16 0 0 1-16-16Zm16 0h64V40H96Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 232a6 6 0 0 1-6 6H56a6 6 0 0 1 0-12h144a6 6 0 0 1 6 6ZM82 192V40a14 14 0 0 1 14-14h64a14 14 0 0 1 14 14v152a14 14 0 0 1-14 14H96a14 14 0 0 1-14-14Zm12 0a2 2 0 0 0 2 2h64a2 2 0 0 0 2-2V40a2 2 0 0 0-2-2H96a2 2 0 0 0-2 2Z"/>`),
		g.Group(children),
	)
}
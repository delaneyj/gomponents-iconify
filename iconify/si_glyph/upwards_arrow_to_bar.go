package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpwardsArrowToBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.002 4.913c.551 0 .995-.42.995-.938V2.102c0-.518-.444-.938-.995-.938H1.076c-.55 0-.995.42-.995.938v1.873c0 .519.445.938.995.938h13.926zM6.994 6.518a1.488 1.488 0 0 1 2.092 0l6.408 6.374c.579.573.835 2.078-.995 2.078H1.579c-1.891 0-1.573-1.505-.995-2.078l6.41-6.374z"/>`),
		g.Group(children),
	)
}
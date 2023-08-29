package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quotedbl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 0h95v113L39 231L3 213L52 95H0V0zm167 0h95v113l-56 118l-37-18l50-118h-52V0z"/>`),
		g.Group(children),
	)
}
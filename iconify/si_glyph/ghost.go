package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m.008 7.91l.016 8.095L2.021 13l1.978 3.005L6.062 13l1.959 3.005L10.08 13l1.93 3.005L14.049 13l1.956 3.005l-.017-8.095C15.989 3.56 12.363.031 7.986.031C3.609.031.008 3.559.008 7.91zm4.508.184a1.563 1.563 0 1 1 0-3.125a1.563 1.563 0 1 1 0 3.125zm6.994-.028a1.584 1.584 0 0 1-1.588-1.579c0-.873.711-1.581 1.588-1.581c.878 0 1.588.708 1.588 1.581c0 .872-.71 1.579-1.588 1.579z"/>`),
		g.Group(children),
	)
}
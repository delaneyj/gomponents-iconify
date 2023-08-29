package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinRightBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.148 15.938c-.573.002-1.034-.314-1.032-.885c0-.568.465-1.031 1.038-1.033l5.353-.008L1.314 1.7A.934.934 0 0 1 1.331.373A.948.948 0 0 1 2.665.355l12.441 12.2l.01-5.347c.002-.57.381-1.033.953-1.034c.285-.002.898.379.898 1.027l-.017 8.721l-8.802.016z"/>`),
		g.Group(children),
	)
}
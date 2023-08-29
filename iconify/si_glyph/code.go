package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.74 11.993a.714.714 0 0 0 .493-.195l3.479-3.316a.644.644 0 0 0 0-.939l-3.479-3.318a.72.72 0 0 0-.984 0a.643.643 0 0 0 0 .938l2.987 2.85l-2.987 2.848a.643.643 0 0 0 0 .938a.713.713 0 0 0 .491.194zm-7.439-.11a.727.727 0 0 1-.497-.19L.287 8.425a.62.62 0 0 1 0-.923l3.517-3.268a.74.74 0 0 1 .995 0a.624.624 0 0 1 0 .924l-3.02 2.805l3.02 2.804a.626.626 0 0 1 0 .926a.736.736 0 0 1-.498.19zm2.384 3.006a.672.672 0 0 0 .503-.51L9.934 1.885a.636.636 0 0 0-.488-.768a.661.661 0 0 0-.77.514L5.93 14.125a.636.636 0 0 0 .755.764z"/>`),
		g.Group(children),
	)
}
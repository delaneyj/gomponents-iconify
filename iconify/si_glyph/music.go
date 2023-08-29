package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.942.751L6.035.042v11.097a4.228 4.228 0 0 0-1.924.299c-1.594.651-2.422 2.217-1.965 3.312c.458 1.098 2.029 1.604 3.621.953c1.224-.5 2.073-1.451 2.184-2.362l-.008-9.817c2.627.798 4.52 1.673 4.52 4.284c0 .927 1.52 1.823 1.52-1.812c-.001-3.092-2.653-4.962-6.041-5.245z"/>`),
		g.Group(children),
	)
}
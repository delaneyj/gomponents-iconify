package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 13H7v-2h10v2zm1-6h-1c-1.631 0-3.065.792-3.977 2H18c1.103 0 2 .897 2 2v2c0 1.103-.897 2-2 2h-4.977c.913 1.208 2.347 2 3.977 2h1a4 4 0 0 0 4-4v-2a4 4 0 0 0-4-4zM2 11v2a4 4 0 0 0 4 4h1c1.63 0 3.065-.792 3.977-2H6c-1.103 0-2-.897-2-2v-2c0-1.103.897-2 2-2h4.977C10.065 7.792 8.631 7 7 7H6a4 4 0 0 0-4 4z"/>`),
		g.Group(children),
	)
}
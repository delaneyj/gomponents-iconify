package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 11l-2 2H7v-2h3zm9.641-3.64L22 5l-1.5-1.5l-17 17L5 22l9-9h3v-2h-1l2-2c1.103 0 2 .897 2 2v2c0 1.103-.897 2-2 2h-4.977c.913 1.208 2.347 2 3.977 2h1a4 4 0 0 0 4-4v-2a3.995 3.995 0 0 0-2.359-3.64zM4.36 16.64L6 15c-1.103 0-2-.897-2-2v-2c0-1.103.897-2 2-2h4.977C10.065 7.792 8.631 7 7 7H6a4 4 0 0 0-4 4v2c0 1.623.971 3.013 2.36 3.64z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOuterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1V1Zm1 1v11h11V2H2Zm6 3H7V4h1v1ZM5 8H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm-3 3H7v-1h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
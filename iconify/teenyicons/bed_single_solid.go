package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedSingleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 15h1v-4h13v4h1V4h-1v3H1V0H0v15Z"/><path fill="currentColor" d="M6 6H2V5h4v1Z"/>`),
		g.Group(children),
	)
}
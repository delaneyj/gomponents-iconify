package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtboardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 0v2h1V0H4Zm6 0v2h1V0h-1ZM2 5H0V4h2v1Zm-2 6h2v-1H0v1Zm15-6h-2V4h2v1Zm-2 6h2v-1h-2v1Zm-9 4v-2h1v2H4Zm6-2v2h1v-2h-1ZM4.5 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Z"/>`),
		g.Group(children),
	)
}
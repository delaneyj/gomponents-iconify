package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendPlaneTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 1.346a.5.5 0 0 1 .241.062l18.462 10.154a.5.5 0 0 1 0 .876L3.741 22.593A.5.5 0 0 1 3 22.154V1.846a.5.5 0 0 1 .5-.5ZM5 4.383V11h5v2H5v6.617L18.85 12L5 4.383Z"/>`),
		g.Group(children),
	)
}
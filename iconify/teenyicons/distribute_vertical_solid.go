package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M15 0H0v1h15V0Zm-1 5H1v5h13V5Zm1 9H0v1h15v-1Z"/>`),
		g.Group(children),
	)
}
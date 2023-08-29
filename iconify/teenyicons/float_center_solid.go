package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatCenterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.5 0A1.5 1.5 0 0 0 4 1.5v4A1.5 1.5 0 0 0 5.5 7h4A1.5 1.5 0 0 0 11 5.5v-4A1.5 1.5 0 0 0 9.5 0h-4ZM0 2h2V1H0v1Zm13 0h2V1h-2v1ZM0 6h2V5H0v1Zm13 0h2V5h-2v1ZM0 10h15V9H0v1Zm0 4h15v-1H0v1Z"/>`),
		g.Group(children),
	)
}
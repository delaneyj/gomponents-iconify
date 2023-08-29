package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v4A1.5 1.5 0 0 0 1.5 7h4A1.5 1.5 0 0 0 7 5.5v-4A1.5 1.5 0 0 0 5.5 0h-4ZM9 2h6V1H9v1Zm0 4h6V5H9v1Zm-9 4h15V9H0v1Zm0 4h15v-1H0v1Z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkSpeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.628 24.08a14.704 14.704 0 0 0-22.673 1.725"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.585a23.808 23.808 0 0 0-39 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 20.584l-21.483 8.464a5.558 5.558 0 1 0 6.212 8.779l.015.02Z"/>`),
		g.Group(children),
	)
}
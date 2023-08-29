package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cvs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.139 15.225l-5.8-5.8a4.648 4.648 0 0 0-6.574 0L24 15.19l-5.765-5.765a4.648 4.648 0 0 0-6.573 0l-5.8 5.8a4.648 4.648 0 0 0 0 6.573l5.764 5.765L24 39.937l18.139-18.139a4.648 4.648 0 0 0 0-6.573Z"/>`),
		g.Group(children),
	)
}
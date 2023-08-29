package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintRoller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 4v8h18V4zm2 2h14v4H7zm17 1v2h1v4.25l-9.281 2.781l-.719.219V19h-2v9h6v-9h-2v-1.25l9.281-2.781l.719-.219V7zm-9 14h2v5h-2z"/>`),
		g.Group(children),
	)
}
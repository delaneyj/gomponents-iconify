package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M576 192v896q0 26-19 45t-45 19t-45-19L19 685Q0 666 0 640t19-45l448-448q19-19 45-19t45 19t19 45z"/>`),
		g.Group(children),
	)
}
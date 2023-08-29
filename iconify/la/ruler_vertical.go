package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 0v32h16V0H8zm2 2h12v3h-7v2h7v2h-4v2h4v2h-7v2h7v2h-4v2h4v2h-7v2h7v2h-4v2h4v3H10V2z"/>`),
		g.Group(children),
	)
}
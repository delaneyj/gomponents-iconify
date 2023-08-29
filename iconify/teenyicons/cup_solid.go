package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 4h1V0H7v4ZM5 2v2H4V2h1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6.5A1.5 1.5 0 0 1 1.5 5h9a1.5 1.5 0 0 1 1.415 1H13.5A1.5 1.5 0 0 1 15 7.5v2a1.5 1.5 0 0 1-1.5 1.5H12v.5A3.5 3.5 0 0 1 8.5 15h-5A3.5 3.5 0 0 1 0 11.5v-5ZM12 10h1.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5H12v3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
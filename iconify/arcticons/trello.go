package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2Zm5 4.81a2.21 2.21 0 0 0-2.22 2.22v21a2.22 2.22 0 0 0 2.22 2.22h6.85a2.22 2.22 0 0 0 2.21-2.22v-21a2.21 2.21 0 0 0-2.21-2.22Zm16.09 0a2.22 2.22 0 0 0-2.22 2.22v11.75a2.23 2.23 0 0 0 2.22 2.22h6.85a2.22 2.22 0 0 0 2.22-2.22V12.53a2.21 2.21 0 0 0-2.22-2.22Z"/>`),
		g.Group(children),
	)
}
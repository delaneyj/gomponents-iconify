package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theswitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.35 4.5a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95ZM17 13.06h14.1a1 1 0 0 1 1 1v19.87a1 1 0 0 1-1 1H17a1 1 0 0 1-1-1V14.07a1 1 0 0 1 1-1.01Z"/>`),
		g.Group(children),
	)
}
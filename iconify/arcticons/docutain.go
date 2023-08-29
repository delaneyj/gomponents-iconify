package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docutain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.57 4.5h2.286A17.062 17.062 0 0 1 40.92 21.563v4.875A17.062 17.062 0 0 1 23.856 43.5H11.081V14.989"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.57 4.5v10.489H11.081L21.57 4.5z"/>`),
		g.Group(children),
	)
}
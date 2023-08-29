package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myexpenses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.5 8.5v9a2 2 0 0 0 2 2h9v18a2 2 0 0 1-2 2h-35a2 2 0 0 1-2-2v-27a2 2 0 0 1 2-2Zm11 11l-11-11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.99 30H20l4-6l-3.99-6H28m13.62-.38V6a1.69 1.69 0 0 0-1.9-1.68L6.5 8.5"/>`),
		g.Group(children),
	)
}
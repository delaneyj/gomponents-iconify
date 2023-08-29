package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoftonenote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.81 18.5v11m7.38 0v-11m-7.38 0l7.38 11M5.5 16v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2h-16a2 2 0 0 0-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.23 14V6.5a2 2 0 0 1 2-2H40.5a2 2 0 0 1 2 2v35a2 2 0 0 1-2 2H15.23a2 2 0 0 1-2-2V34"/>`),
		g.Group(children),
	)
}
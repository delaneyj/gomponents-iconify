package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pismowitewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a1.95 1.95 0 0 0 1.95 1.95h2.376v-39H10.35A1.95 1.95 0 0 0 8.4 6.45Zm4.326-1.95v39H37.65a1.95 1.95 0 0 0 1.95-1.95V6.45a1.95 1.95 0 0 0-1.95-1.95Zm13.437 23.626v-17m-5 5.192h10"/>`),
		g.Group(children),
	)
}
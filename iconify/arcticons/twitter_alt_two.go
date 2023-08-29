package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.5 41.5h-7.368L6.5 6.5h7.368l27.632 35zM21.941 26.059L6.5 41.5m35-35L26.059 21.941"/>`),
		g.Group(children),
	)
}
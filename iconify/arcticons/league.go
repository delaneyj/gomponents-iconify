package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func League(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.844 4.5h11.103v31.443h15.209L35.103 43.5H11.871l4.198-7.557V12.337L10.844 4.5zm3.072 4.608v30.71"/>`),
		g.Group(children),
	)
}
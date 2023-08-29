package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 0H0v15h15V2.5a.5.5 0 0 0-.907-.29L9.49 8.653L6.9 5.2a.5.5 0 0 0-.816.023L1 12.849V0Z"/>`),
		g.Group(children),
	)
}
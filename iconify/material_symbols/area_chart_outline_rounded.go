package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChartOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.425 0-.713-.288T3 19V7.25q.125 0 .288.038t.312.162L7 10l4.375-6.15q.25-.35.675-.413t.775.213L17 7h3q.425 0 .713.288T21 8v11q0 .425-.288.713T20 20H4Zm4-3l3.4-4.675q.25-.35.663-.413t.762.213L19 16.95V9h-2.7l-3.9-3.125l-4.95 6.95L5 11v3.6L8 17Z"/>`),
		g.Group(children),
	)
}
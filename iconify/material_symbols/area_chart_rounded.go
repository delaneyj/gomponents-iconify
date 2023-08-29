package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 16l-7.775-6.075q-.675-.525-1.512-.413t-1.338.813l-2.75 3.8L4.25 11.45q-.3-.25-.625-.35T3 11V7.25q.125 0 .288.038t.312.162L7 10l4.375-6.15q.25-.35.675-.413t.775.213L17 7h3q.425 0 .713.288T21 8v8ZM4 20q-.425 0-.713-.288T3 19v-5.725q.175 0 .325.05t.3.175L8 17l3.4-4.675q.25-.35.663-.413t.762.213l8.175 6.4V19q0 .425-.288.713T20 20H4Z"/>`),
		g.Group(children),
	)
}
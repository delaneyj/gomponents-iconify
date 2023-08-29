package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.075 21q-.4 0-.762-.163t-.638-.437l-4.1-4.125Q3.3 16 3.15 15.638T3 14.875v-5.8q0-.4.15-.762t.425-.638L4.2 7.05L1.4 4.2q-.275-.275-.287-.688T1.4 2.8q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.85-2.85l-.6.625q-.275.3-.65.463t-.8.162H9.075ZM12 17q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm7.8-.05l-6.8-6.8V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v.15L7.05 4.2l.625-.625q.275-.275.638-.425T9.075 3h5.85q.4 0 .763.15t.637.425l4.1 4.1q.275.275.425.638t.15.762V14.9q0 .4-.138.75t-.412.625l-.65.675Z"/>`),
		g.Group(children),
	)
}
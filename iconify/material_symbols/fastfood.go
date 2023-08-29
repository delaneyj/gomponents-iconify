package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastfood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 15q0-1.5.65-2.625t1.7-1.875q1.05-.75 2.4-1.125T8.5 9q1.4 0 2.75.375t2.4 1.125q1.05.75 1.7 1.875T16 15H1Zm17 8v-8q0-2.875-1.763-4.888T11.276 7.3L11 5h5V1h2v4h5l-1.65 16.55q-.075.6-.537 1.025T19.7 23H18ZM1 19v-2h15v2H1Zm1 4q-.425 0-.713-.288T1 22v-1h15v1q0 .425-.288.713T15 23H2Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodtypeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.425 0-5.713-2.35T4 13.8q0-1.55.7-3.1t1.75-2.975Q7.5 6.3 8.725 5.05T11 2.875q.2-.2.463-.287T12 2.5q.275 0 .537.088t.463.287q1.05.925 2.275 2.175t2.275 2.675Q18.6 9.15 19.3 10.7t.7 3.1q0 3.5-2.288 5.85T12 22Zm0-2q2.6 0 4.3-1.763T18 13.8q0-1.825-1.513-4.125T12 4.65Q9.025 7.375 7.513 9.675T6 13.8q0 2.675 1.7 4.438T12 20Zm0-8Zm-2 6h4q.425 0 .713-.288T15 17q0-.425-.288-.713T14 16h-4q-.425 0-.713.288T9 17q0 .425.288.713T10 18Zm1-5v1q0 .425.288.713T12 15q.425 0 .713-.288T13 14v-1h1q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11h-1v-1q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10v1h-1q-.425 0-.713.288T9 12q0 .425.288.713T10 13h1Z"/>`),
		g.Group(children),
	)
}
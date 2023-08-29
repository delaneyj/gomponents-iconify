package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterEcRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.425 0-5.713-2.35T4 13.8q0-1.55.7-3.1t1.75-2.975Q7.5 6.3 8.725 5.05T11 2.875q.2-.2.463-.287T12 2.5q.275 0 .537.088t.463.287q1.05.925 2.275 2.175t2.275 2.675Q18.6 9.15 19.3 10.7t.7 3.1q0 3.5-2.288 5.85T12 22Zm-1-8v3.85q0 .375.363.475t.562-.2l3.55-5.35q.175-.25.038-.512T15.075 12H13V8.15q0-.375-.363-.475t-.562.2l-3.55 5.35q-.175.25-.038.513t.438.262H11Z"/>`),
		g.Group(children),
	)
}
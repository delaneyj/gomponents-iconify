package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHomeWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7h2V5h-2v2Zm1 14q-2.075 0-3.538-1.463T13 16q0-2.075 1.463-3.538T18 11q2.075 0 3.538 1.463T23 16q0 2.075-1.463 3.538T18 21Zm-.5-2h1v-2.5H21v-1h-2.5V13h-1v2.5H15v1h2.5V19Zm5.5-7.875q-.975-1.05-2.275-1.588T18 9q-.275 0-.513.013T17 9.074V8l-7-5.05V1h13v10.125ZM1 19V9l7-5l7 5v.675q-1.8.85-2.9 2.588T11 16q0 .775.163 1.538T11.675 19H10v-6H6v6H1Z"/>`),
		g.Group(children),
	)
}
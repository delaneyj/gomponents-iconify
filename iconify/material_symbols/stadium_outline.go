package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StadiumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7V3l4 2l-4 2Zm15 0V3l4 2l-4 2Zm-7-1V2l4 2l-4 2Zm0 16q-1.9-.05-3.538-.313t-2.85-.662Q3.4 20.625 2.7 20.1T2 19v-9q0-.625.788-1.163t2.137-.95q1.35-.412 3.175-.65T12 7q2.075 0 3.9.238t3.175.65q1.35.412 2.138.95T22 10v9q0 .575-.7 1.1t-1.913.925q-1.212.4-2.85.663T13 22v-4h-2v4Zm1-11q2.425 0 4.188-.288T19 10.05q0-.125-1.9-.588T12 9q-3.2 0-5.1.463T5 10.05q1.05.375 2.813.663T12 11Zm-3 8.85V16h6v3.85q2-.2 3.275-.588T20 18.576V11.8q-1.375.55-3.45.875T12 13q-2.475 0-4.55-.325T4 11.8v6.775q.45.3 1.725.688T9 19.85Zm3-4.025Z"/>`),
		g.Group(children),
	)
}
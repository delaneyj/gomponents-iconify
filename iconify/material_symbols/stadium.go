package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stadium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7V3l4 2l-4 2Zm15 0V3l4 2l-4 2Zm-7-1V2l4 2l-4 2ZM9 21.875q-1.3-.125-2.55-.375t-2.238-.613q-.987-.362-1.6-.837T2 19v-9q0-.625.788-1.163t2.137-.95q1.35-.412 3.175-.65T12 7q2.075 0 3.9.238t3.175.65q1.35.412 2.138.95T22 10v9q0 .575-.613 1.05t-1.6.838q-.987.362-2.237.612t-2.55.375V17H9v4.875ZM12 11q2.425 0 4.188-.288T19 10.05q0-.125-1.9-.588T12 9q-3.2 0-5.1.463T5 10.05q1.05.375 2.813.663T12 11Z"/>`),
		g.Group(children),
	)
}
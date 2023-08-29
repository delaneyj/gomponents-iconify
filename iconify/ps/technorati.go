package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Technorati(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M77 8q12-2 36-2h29q-40 51-44 95q-5 40 13 78.5t51 65.5q4 2 6 6t2 7.5t-.5 9.5t-.5 8q-7 65-8 79q13-8 45.5-30.5T258 292q117 22 204-33v91q1 26-13.5 49.5T410 433q-17 9-48 9H225l-137-1q-34 0-60-26T2 354V101Q0 69 22.5 41.5T77 8zm120 59q-26 28-26 60q-3 46 39 78q37 28 87.5 30.5T389 215q55-31 55-84q2-48-42-79q-42-31-100-29q-64 2-105 44z"/>`),
		g.Group(children),
	)
}
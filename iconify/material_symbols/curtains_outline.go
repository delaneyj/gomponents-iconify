package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurtainsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2H2Zm4-2h3.95q-.2-1.75-1.15-3.538T6 13.15V19ZM6 5v5.85q1.85-.525 2.8-2.312T9.95 5H6Zm2.225 7q1.7 1.125 2.625 3.075T11.95 19h.1q.175-1.975 1.1-3.925T15.775 12q-1.7-1.125-2.625-3.075T12.05 5h-.1q-.175 1.975-1.1 3.925T8.225 12ZM18 5h-3.95q.2 1.75 1.15 3.538T18 10.85V5Zm0 14v-5.85q-1.85.525-2.788 2.313T14.075 19H18ZM6 5v5.85V5Zm12 0v5.85V5Zm0 14v-5.85V19ZM6 19v-5.85V19Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWifiProOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q.625 0 1.063-.438T13.5 15.5q0-.625-.438-1.063T12 14q-.625 0-1.063.438T10.5 15.5q0 .625.438 1.063T12 17Zm-3 5q-1.925 0-3.1-1t-1.825-2.5q-.65-1.525-.863-3.25T3 12.025q0-2.55.388-4.413T4.8 4.476q1.025-1.25 2.75-1.862T12 2q2.725 0 4.45.613t2.75 1.862q1.025 1.275 1.413 3.138T21 12.024q0 1.5-.213 3.225t-.862 3.25Q19.275 20 18.1 21T15 22H9Zm-.025-2h6q1.35 0 2.125-.875t1.2-2.15q.425-1.275.55-2.638t.125-2.312q0-2.125-.25-3.637T17.7 5.9q-.75-.975-2.113-1.438T11.975 4q-2.25 0-3.612.463T6.25 5.9q-.775.975-1.025 2.488t-.25 3.637q0 .95.125 2.313t.55 2.637q.425 1.275 1.2 2.15T8.975 20Zm3-8Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainVerification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4ZM4 8h16V6H4v2Zm6.95 8.55l5.65-5.65l-1.45-1.45l-4.2 4.2l-2.1-2.1L7.4 13l3.55 3.55Z"/>`),
		g.Group(children),
	)
}
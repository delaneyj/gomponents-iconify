package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditScoreRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6.275q-.85-.45-1.838-.325t-1.712.85l-3.5 3.525l-.75-.725q-.875-.875-2.1-.863t-2.1.888q-.9.9-.9 2.125t.875 2.1l.15.15H4ZM4 8v4h16V8H4Zm10.95 11.15l4.925-4.925q.3-.3.713-.288t.712.313q.275.3.288.7t-.288.7l-5.65 5.65q-.3.3-.7.3t-.7-.3l-2.85-2.85q-.275-.275-.288-.688t.288-.712q.275-.275.688-.275t.712.275l2.15 2.1Z"/>`),
		g.Group(children),
	)
}
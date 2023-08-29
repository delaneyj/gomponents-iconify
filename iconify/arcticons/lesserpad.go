package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lesserpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.6 18.79V6.45a2 2 0 0 0-2-1.95H10.35a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V24.49M15.22 12.7h17.55m0 7.53H15.22m0 7.54h9.75m-9.75 7.53h5.85"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26 32.84v2.46h2.4a1 1 0 0 0 .68-.29L44.59 19.5l-2.85-2.85l-15.51 15.51a1 1 0 0 0-.23.68Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M46.78 17.29a.75.75 0 0 0 0-1.06h0l-1.73-1.77a.75.75 0 0 0-1.06 0l-2.25 2.2l2.85 2.85Z"/>`),
		g.Group(children),
	)
}
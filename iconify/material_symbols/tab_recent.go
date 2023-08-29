package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabRecent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.65 20.35l.7-.7l-1.85-1.85V15h-1v3.2l2.15 2.15ZM13 10h7V6h-7v4Zm5 13q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6.275q-.875-.625-1.913-.95T17.975 11q-2.9 0-4.938 2.05T11 18q0 .525.075 1.025T11.3 20H4Z"/>`),
		g.Group(children),
	)
}
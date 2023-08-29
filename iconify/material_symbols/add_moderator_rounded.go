package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddModeratorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 17.5v2q0 .2.15.35T17 20q.2 0 .35-.15t.15-.35v-2h2q.2 0 .35-.15T20 17q0-.2-.15-.35t-.35-.15h-2v-2q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35v2h-2q-.2 0-.35.15T14 17q0 .2.15.35t.35.15h2ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm-5 0q-3.475-.875-5.738-3.988T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375v4.3q-.65-.325-1.463-.5T17 10q-2.9 0-4.95 2.05T10 17q0 1.55.588 2.8t1.487 2.175q-.025 0-.037.013T12 22Z"/>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserMilitary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25 13l-1.593 3l-3.407.414l2.5 2.253L22 22l3-1.875L28 22l-.5-3.333l2.5-2.253L26.5 16L25 13z"/><path fill="currentColor" d="M21.414 13.414L25 9.834l3.587 3.582L30 12l-5-5l-5 5l1.414 1.414z"/><path fill="currentColor" d="M21.414 8.414L25 4.834l3.587 3.582L30 7l-5-5l-5 5l1.414 1.414zM16 30h-2v-5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v5H2v-5a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM9 10a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z"/>`),
		g.Group(children),
	)
}
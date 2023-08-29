package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Splitwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 5.5l17.249 9.959V42.5H6.75V15.474Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.065 40.763a4.433 4.433 0 0 0 3.886 1.737h2.347a3.958 3.958 0 0 0 3.954-3.963h0a3.958 3.958 0 0 0-3.954-3.962h-2.593a3.958 3.958 0 0 1-3.954-3.963h0a3.958 3.958 0 0 1 3.954-3.962h2.347a4.433 4.433 0 0 1 3.886 1.736m24.311-12.927l-24.31 12.927M6.751 15.474L41.25 33.819"/>`),
		g.Group(children),
	)
}
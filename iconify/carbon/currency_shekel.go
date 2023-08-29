package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyShekel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 27H7V5h8a5.006 5.006 0 0 1 5 5v12h-2V10a3.003 3.003 0 0 0-3-3H9Z"/><path fill="currentColor" d="M20 27h-8V10h2v15h6a3.003 3.003 0 0 0 3-3V5h2v17a5.006 5.006 0 0 1-5 5Z"/>`),
		g.Group(children),
	)
}
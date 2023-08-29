package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBaht(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.402 15.643A4.973 4.973 0 0 0 23 12v-1a5.006 5.006 0 0 0-5-5h-1V3h-2v3h-5v20h5v3h2v-3h2a5.006 5.006 0 0 0 5-5v-1a4.983 4.983 0 0 0-2.598-4.357ZM12 8h6a3.003 3.003 0 0 1 3 3v1a3.003 3.003 0 0 1-3 3h-6Zm10 13a3.003 3.003 0 0 1-3 3h-7v-7h7a3.003 3.003 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}
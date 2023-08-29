package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quoteunquote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 29.75c2.21 0 4-2.686 4-6h-4v-9.5h8.7v9.5c0 5.523-3.895 10-8.7 10Zm13.3 0c2.21 0 4-2.686 4-6h-4v-9.5H35v9.5c0 5.523-3.895 10-8.7 10Z"/>`),
		g.Group(children),
	)
}
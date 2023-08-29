package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freetrailer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.921 19.649h11.86v-5.1l9.43 9.43l-9.43 9.43v-5.1H2.922m42.08.191H33.328l-.18 4.95l-9.01-9.54l9.16-9.34v5.27h11.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.921 19.649a21.505 21.505 0 0 1 42.157.192m-.076 8.66a21.495 21.495 0 0 1-42.08-.192"/>`),
		g.Group(children),
	)
}
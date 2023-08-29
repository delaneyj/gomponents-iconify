package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xpa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#4FA784"/><path fill="#FFF" d="m25.575 12.846l-8.11-2.324l-2.774 4.725l-3.222-3.737l1.46-.188l1.505 1.768l1.67-2.959l-7.35-2.107a.586.586 0 0 0-.736.7L11.274 21.8l3.497-5.888l3.222 3.737l-1.46.188l-1.504-1.768l-3.31 5.517l.49 1.97a.59.59 0 0 0 .98.28l12.63-12.01a.58.58 0 0 0-.244-.98z"/></g>`),
		g.Group(children),
	)
}
package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xpa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm9.575-19.154l-8.11-2.324l-2.774 4.725l-3.222-3.737l1.46-.188l1.505 1.768l1.67-2.959l-7.35-2.107a.586.586 0 0 0-.736.7L11.274 21.8l3.497-5.888l3.222 3.737l-1.46.188l-1.504-1.768l-3.31 5.517l.49 1.97a.59.59 0 0 0 .98.28l12.63-12.01a.58.58 0 0 0-.244-.98z"/>`),
		g.Group(children),
	)
}
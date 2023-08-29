package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassengerDrinks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 4V2h-8v2a3.99 3.99 0 0 0 3 3.859V11h-2v2h6v-2h-2V7.859A3.99 3.99 0 0 0 30 4zM10.5 9A3.5 3.5 0 1 1 14 5.5A3.504 3.504 0 0 1 10.5 9zm0-5A1.5 1.5 0 1 0 12 5.5A1.502 1.502 0 0 0 10.5 4zm11.975 27.313L19.34 24h-7.102a4.007 4.007 0 0 1-3.866-2.97l-1.634-6.127a3.899 3.899 0 0 1 7.534-2.009l.83 3.106H21v2h-7.435l-1.223-4.59a1.9 1.9 0 0 0-3.671.978l1.634 6.126A2.005 2.005 0 0 0 12.239 22h8.42l3.655 8.525z"/><path fill="currentColor" d="M18 28H7.768a2.002 2.002 0 0 1-1.932-1.485L2.034 12.258l1.932-.516L7.768 26H18Z"/>`),
		g.Group(children),
	)
}
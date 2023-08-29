package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobilbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.287 25.025s.88 4.083 4.801 4.083s4.456-4.097 4.456-4.097m-7.288 6.323s.849 7.894 1.514 10.657m1.533 1.509s1.525-3.835 2.201-12.004"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.617 17.317c-2.459 1.847-3.39 3.191-2.482 6.309s1.268 3.19 1.268 3.19a5.587 5.587 0 0 1 2.025-5.27c2.683-2.092 3.994-4.781 2.788-8.827s-4.7-8.22-4.7-8.22s.294 5.749-2.281 8.478c-1.517 1.608-5 3.624-3.56 7.501s2.725 3.288 2.725 3.288s-2.274-3.908.591-6.389s4.217-3.396 4.278-5.942c0 0 1.806 4.035-.652 5.882Z"/>`),
		g.Group(children),
	)
}
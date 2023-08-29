package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AarogyaSetu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.187 27.686C32.695 33.671 24 41.703 24 41.703s-3.505-3.238-7.47-7.134m-3.197-3.193c-3.345-3.402-6.446-6.787-7.515-8.638c-2.684-4.648-1.23-11.945 4.011-14.97S21.774 6.536 24 11.777c2.226-5.241 8.929-7.037 14.17-4.01s6.696 10.322 4.012 14.97a14.104 14.104 0 0 1-1.013 1.47"/>`),
		g.Group(children),
	)
}
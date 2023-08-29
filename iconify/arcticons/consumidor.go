package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Consumidor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.073 29.811V18.19h3.78a3.922 3.922 0 1 1 0 7.845h-3.78M9.645 28.503a3.467 3.467 0 0 0 2.908 1.307h1.745a2.905 2.905 0 1 0 0-5.81h-1.89a2.905 2.905 0 1 1 0-5.812h1.745a3.123 3.123 0 0 1 2.908 1.308M38.5 25.961a3.926 3.926 0 0 1-7.851 0V22.04a3.908 3.908 0 0 1 3.925-3.922a3.787 3.787 0 0 1 3.78 3.922"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
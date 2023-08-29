package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Todotree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.011 12.763h7.674v7.674h-7.674zM5.5 5.637V16.6h11.511m3.837 3.838v7.126h13.705m0 7.124h7.674v7.674h-7.674z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.848 27.564v10.963h13.705m1.917-12.334l1.645 2.866l4.385-9.444"/><path d="M38.477 23.727h-3.925V31.4h7.674v-7.675"/></g>`),
		g.Group(children),
	)
}
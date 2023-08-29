package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zomato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 21.51h3.306L9 25.89h3.306m3.306 0h0a1.658 1.658 0 0 1-1.653-1.653v-1.074a1.658 1.658 0 0 1 1.653-1.653h0a1.658 1.658 0 0 1 1.652 1.653v1.074a1.658 1.658 0 0 1-1.652 1.653Zm3.305-2.727a1.658 1.658 0 0 1 1.653-1.653h0a1.658 1.658 0 0 1 1.653 1.653v2.645m-3.306-4.298v4.298m3.306-2.645a1.658 1.658 0 0 1 1.653-1.653h0a1.658 1.658 0 0 1 1.653 1.653v2.645m4.959-1.571a1.658 1.658 0 0 1-1.653 1.653h0a1.658 1.658 0 0 1-1.653-1.653v-1.074a1.658 1.658 0 0 1 1.653-1.653h0a1.658 1.658 0 0 1 1.653 1.653m0 2.727v-4.38m2.479-1.404v4.959a.78.78 0 0 0 .826.826h.248m-1.9-4.363h1.735m3.471 4.362h0a1.658 1.658 0 0 1-1.653-1.653v-1.074a1.658 1.658 0 0 1 1.653-1.653h0A1.658 1.658 0 0 1 39 23.163v1.074a1.658 1.658 0 0 1-1.653 1.653Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}
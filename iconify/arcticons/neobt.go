package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neobt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.329 5.76c-.53 5.359-1.423 11.164-1.423 11.164c-5.415-.626-9.3-.139-9.3-.139c-.012 3.41-.789 8.326-.789 8.326L24 31.235s2.806-2.149 4.152-4.406c0 0-.484-3.363-.565-5.996l-2.818-.093l.035-11.19c6.463-.077 10.715 1.268 14.774 2.435c-1.624 9.185-.464 12.903.549 15.95c-2.845 7.89-9.618 13.341-16.127 16.472c-7.105-2.876-12.973-10.646-16.127-18.253c1.113-3.038 2.32-8.86.58-18.09C21.209 3.913 31.329 5.76 31.329 5.76Z"/>`),
		g.Group(children),
	)
}
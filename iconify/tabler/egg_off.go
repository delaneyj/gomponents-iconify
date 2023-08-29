package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.927 17.934C16.716 19.792 14.576 20.887 12 21c-4.2 0-7-2.763-7-6.917c0-2.568.753-5.14 1.91-7.158m1.732-2.297C9.676 3.608 10.838 2.998 12 3c3.5.007 7 5.545 7 11.083c0 .298-.015.587-.045.868M3 3l18 18"/>`),
		g.Group(children),
	)
}
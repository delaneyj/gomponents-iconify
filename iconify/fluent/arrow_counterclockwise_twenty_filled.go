package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterclockwiseTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.372 2.025A8 8 0 1 1 2.005 9.74a.75.75 0 0 1 1.5.045a6.5 6.5 0 1 0 1.573-4.029l-.204.248h2.379l.102.008a.75.75 0 0 1 0 1.486l-.102.007h-4l-.102-.007a.75.75 0 0 1-.641-.641l-.007-.102v-4l.007-.102a.75.75 0 0 1 .641-.641l.102-.007l.102.007a.75.75 0 0 1 .641.641l.007.102l-.001 1.953a7.977 7.977 0 0 1 5.37-2.682Z"/>`),
		g.Group(children),
	)
}
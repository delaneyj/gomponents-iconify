package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quickedit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.82 2.548a21.321 21.321 0 1 0 9.298 40.512l-5.903-5.902a13.728 13.728 0 1 1 9.935-10.045l5.924 5.924a21.33 21.33 0 0 0-19.253-30.49ZM43.753 37.13a23.69 23.69 0 0 1-6.556 6.595l1.727 1.728H45.5v-6.577Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.657 25.034l-6.576 6.576l2.615 2.614a11.583 11.583 0 0 0 6.532-6.62Zm-3.961 9.19l9.5 9.5m6.557-6.594l-9.525-9.526"/>`),
		g.Group(children),
	)
}
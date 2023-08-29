package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.908 12.917a5 5 0 1 0-5.827-5.819m-.965 3.027l-6.529 7.46a2 2 0 1 0 2.827 2.83l7.461-6.529M3 3l18 18"/>`),
		g.Group(children),
	)
}
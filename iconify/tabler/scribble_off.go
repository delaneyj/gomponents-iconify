package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScribbleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15c2 3 4 4 7 4c1.95 0 4.324-1.268 5.746-3.256m1.181-2.812A5.97 5.97 0 0 0 17 12c0-4-3-7-6-7c-.642 0-1.239.069-1.78.201M6.728 6.716C6.258 7.333 6 8.102 6 9c0 2.5 2 5 6 5c.597 0 1.203-.055 1.808-.156m3.102-.921C19.145 11.97 21.062 10.5 22 9M3 3l18 18"/>`),
		g.Group(children),
	)
}
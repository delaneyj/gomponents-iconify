package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.174 5A6.503 6.503 0 0 1 13.78 2.708l-.812.584A5.502 5.502 0 0 0 3.207 5H7v1H3.022A5.57 5.57 0 0 0 3 6.5v2c0 .169.008.335.022.5H7v1H3.207a5.502 5.502 0 0 0 9.761 1.708l.812.584A6.503 6.503 0 0 1 2.174 10H0V9h2.019A6.593 6.593 0 0 1 2 8.5v-2c0-.168.006-.335.019-.5H0V5h2.174Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
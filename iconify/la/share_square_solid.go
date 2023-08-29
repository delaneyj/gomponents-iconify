package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareSquareSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.414 4.438L22 5.851L26.168 10H16.5a5.506 5.506 0 0 0-5.5 5.5c0 3.033 2.468 5.5 5.5 5.5h.5v-2h-.5c-1.93 0-3.5-1.57-3.5-3.5s1.57-3.5 3.5-3.5h9.672l-4.164 4.164l1.414 1.414L30 11l-6.586-6.563zM5 5v22h22V17l-2 2v6H7V7h10.854l2-2H5z"/>`),
		g.Group(children),
	)
}
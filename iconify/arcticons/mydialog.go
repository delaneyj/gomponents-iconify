package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mydialog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.07 12.049V43.5M10.639 19.911l6.716-3.93l6.715-3.932l6.714 3.932l6.716 3.93v15.726l-6.716 3.931L24.07 43.5l-6.715-3.932l-6.716-3.931M37.5 19.911L10.639 35.637m0-15.726L37.5 35.637M24.07 12.049L10.5 4.5l.139 15.411"/>`),
		g.Group(children),
	)
}
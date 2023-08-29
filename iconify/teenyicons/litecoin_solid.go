package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LitecoinSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m3.714 6.584l1.3-5.205l.971.242l-1.093 4.374l1.884-.942l.448.894l-2.652 1.326L3.14 13H13v1H1.86l1.534-6.138l-2.17 1.085l-.448-.894l2.938-1.469Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
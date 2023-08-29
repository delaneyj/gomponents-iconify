package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10.854.146a.5.5 0 0 0-.708 0L0 10.293V14.5a.5.5 0 0 0 .5.5h4.207L14.854 4.854a.5.5 0 0 0 0-.708l-4-4Z"/>`),
		g.Group(children),
	)
}
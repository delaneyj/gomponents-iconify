package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nurse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6c2.941 0 5.685.847 8 2.31L18 18H6L4 8.309A14.93 14.93 0 0 1 12 6zm-2 6h4m-2-2v4"/>`),
		g.Group(children),
	)
}
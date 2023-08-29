package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderGenderqueer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11a5 5 0 1 1 0 10a5 5 0 0 1 0-10zm0 0V3m2.5 1.5l-5 3m0-3l5 3"/>`),
		g.Group(children),
	)
}
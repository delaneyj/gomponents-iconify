package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.906 6.594l-.718.687l-3.907 3.907l-.686.72l.687.687l13 13l.72.718l.72-.718l13-13l.686-.688l-.687-.72l-3.907-3.905l-.72-.686l-.687.687L16 15.688L7.594 7.28l-.688-.686zm-.03 2.843l8.405 8.376l.72.687l.72-.688l8.405-8.375l2.438 2.438L16 23.47L4.437 11.874l2.438-2.438z"/>`),
		g.Group(children),
	)
}
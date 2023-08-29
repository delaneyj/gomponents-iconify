package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.243 6.343L14.828 4.93L7.758 12l7.07 7.071l1.415-1.414L10.586 12l5.657-5.657Z"/>`),
		g.Group(children),
	)
}
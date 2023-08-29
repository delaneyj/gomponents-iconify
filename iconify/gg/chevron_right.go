package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.586 6.343L12 4.93L19.071 12L12 19.071l-1.414-1.414L16.243 12l-5.657-5.657Z"/>`),
		g.Group(children),
	)
}
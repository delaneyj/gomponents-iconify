package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.53 6.47a.75.75 0 0 0-1.06 0l-5 5a.75.75 0 0 0 0 1.06l5 5a.75.75 0 0 0 1.06 0a.75.75 0 0 0 0-1.06L5.06 12l4.47-4.47a.75.75 0 0 0 0-1.06Zm11 5l-5-5a.75.75 0 0 0-1.06 1.06L18.94 12l-4.47 4.47a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0l5-5a.75.75 0 0 0 0-1.06Z"/>`),
		g.Group(children),
	)
}
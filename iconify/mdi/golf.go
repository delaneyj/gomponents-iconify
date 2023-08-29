package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Golf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 18a1.5 1.5 0 0 1 1.5 1.5a1.5 1.5 0 0 1-1.5 1.5a1.5 1.5 0 0 1-1.5-1.5a1.5 1.5 0 0 1 1.5-1.5M17 5.92L11 9v9.03c2.84.16 5 .97 5 1.97c0 1.1-2.69 2-6 2s-6-.9-6-2c0-.74 1.21-1.38 3-1.73V20h2V2l8 3.92Z"/>`),
		g.Group(children),
	)
}
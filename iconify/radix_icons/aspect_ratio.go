package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AspectRatio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 2h10a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.5-.5v-10a.5.5 0 0 1 .5-.5ZM1 2.5A1.5 1.5 0 0 1 2.5 1h10A1.5 1.5 0 0 1 14 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-10ZM7.5 4a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM8 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
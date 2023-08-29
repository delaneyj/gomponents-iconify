package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoPetsAllowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l9.082 9.082M23.5 23.5L9.582 9.582M16 22v-3m0-3V7.5h1A2.5 2.5 0 0 0 19.5 5V3.5H14A2.5 2.5 0 0 1 11.5 1v3.894A6.74 6.74 0 0 1 9.68 9.5l-.098.082M7.5 11.315a6.737 6.737 0 0 0-2 4.79V21.5H5A2.5 2.5 0 0 1 2.5 19v-4m5.5.5h2.5v.5c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5"/>`),
		g.Group(children),
	)
}
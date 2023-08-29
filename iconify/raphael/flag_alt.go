package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.562 10.75C21.74 8.572 25.5 7 25.5 7c-8 0-8-4-16-4v10c8 0 8 4 16 4c0 0-3.75-3-5.938-6.25zM6.5 29h2V3h-2v26z"/>`),
		g.Group(children),
	)
}
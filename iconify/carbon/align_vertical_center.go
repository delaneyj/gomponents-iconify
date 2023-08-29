package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15h-4v-3a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v3h-4V8a2.002 2.002 0 0 0-2-2H8a2.002 2.002 0 0 0-2 2v7H2v2h4v7a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2v-7h4v3a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2v-3h4ZM8 24V8h4l.001 16Zm12-4v-8h4l.001 8Z"/>`),
		g.Group(children),
	)
}
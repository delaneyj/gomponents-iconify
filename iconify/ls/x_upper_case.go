package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M535 0L316 357l232 378h-84L274 427L85 735H0l231-378L12 0h84l178 289L451 0h84z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Treecard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="25.54" height="32.36" x="11.23" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.05" ry="4.05"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5v39M11.23 10.19C11.23 20.94 24 17.62 24 27.7m12.77-9.07C36.77 29.38 24 26.06 24 36.14M15.94 8.89c-.33 0-.65.04-.96.11c-.07.31-.11.63-.11.96c0 2.44 1.98 4.42 4.42 4.42c.33 0 .65-.04.96-.11c.07-.31.11-.63.11-.96c0-2.44-1.98-4.42-4.42-4.42Zm16.12 4.57c.33 0 .65.04.96.11c.07.31.11.63.11.96c0 2.44-1.98 4.42-4.42 4.42c-.33 0-.65-.04-.96-.11c-.07-.31-.11-.63-.11-.96c0-2.44 1.98-4.42 4.42-4.42ZM15.94 25.29c-.33 0-.65.04-.96.11c-.07.31-.11.63-.11.96c0 2.44 1.98 4.42 4.42 4.42c.33 0 .65-.04.96-.11c.07-.31.11-.63.11-.96c0-2.44-1.98-4.42-4.42-4.42Z"/>`),
		g.Group(children),
	)
}
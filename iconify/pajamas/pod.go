package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.248 1.87l4.485 2.562L8 7.136L3.268 4.432l4.484-2.563a.5.5 0 0 1 .496 0ZM2.5 10.84V5.72l4.75 2.714v5.409l-4.498-2.57a.5.5 0 0 1-.252-.435Zm6.25 3.004l4.498-2.57a.5.5 0 0 0 .252-.435V5.721L8.75 8.435v5.409ZM8.992.567a2 2 0 0 0-1.984 0l-5 2.857A2 2 0 0 0 1 5.161v5.678a2 2 0 0 0 1.008 1.737l5 2.857a2 2 0 0 0 1.984 0l5-2.857A2 2 0 0 0 15 10.839V5.161a2 2 0 0 0-1.008-1.737l-5-2.857Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
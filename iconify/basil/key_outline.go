package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/><path fill="currentColor" fill-rule="evenodd" d="M7 6.25a5.75 5.75 0 1 0 5.05 8.5h3.2V17c0 .414.336.75.75.75h3.5a.75.75 0 0 0 .75-.75v-2.25H22a.75.75 0 0 0 .75-.75v-3A1.75 1.75 0 0 0 21 9.25h-8.95a5.749 5.749 0 0 0-5.05-3ZM2.75 12a4.25 4.25 0 0 1 8.147-1.7a.75.75 0 0 0 .687.45H21a.25.25 0 0 1 .25.25v2.25H19.5a.75.75 0 0 0-.75.75v2.25h-2V14a.75.75 0 0 0-.75-.75h-4.416a.75.75 0 0 0-.687.45A4.251 4.251 0 0 1 2.75 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
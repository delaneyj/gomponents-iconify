package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h4.25a.75.75 0 0 0 0-1.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v4.25a.75.75 0 0 0 1.5 0V3a2 2 0 0 0-2-2H3Zm12 9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5h-2.25a.75.75 0 0 0 0 1.5H13a2 2 0 0 0 2-2v-2.25Zm-3.867-.883L9.5 8l-2.476 2.83L5.5 9L4 10.8V12h5l2.133-2.133ZM6.5 8a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
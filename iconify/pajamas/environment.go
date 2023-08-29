package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Environment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.712 5.997l3.291-3.29l3.29 3.29l-3.29 3.29l-3.29-3.29Zm2.584-4.705a1 1 0 0 1 1.414 0l3.998 3.998a1 1 0 0 1 0 1.414l-3.998 3.998a1 1 0 0 1-1.414 0L7.827 9.233l-1.242 1.243a3 3 0 1 1-1.06-1.06l1.242-1.243l-1.469-1.469a1 1 0 0 1 0-1.414l3.998-3.998ZM4 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
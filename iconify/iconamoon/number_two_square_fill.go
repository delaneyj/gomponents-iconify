package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14Zm7.94 4.438a1.5 1.5 0 0 1 2.12 2.122l-3.828 3.828a2.5 2.5 0 0 0-.732 1.768V17a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-3.975a.5.5 0 0 1 .121-.197l3.829-3.828A3.5 3.5 0 1 0 8.5 9.5a1 1 0 1 0 2 0c0-.385.146-.768.44-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
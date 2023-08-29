package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 15H20V7a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v8H2.5a.5.5 0 0 0-.5.5V17a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3v-1.5a.5.5 0 0 0-.5-.5zM5 7a2.003 2.003 0 0 1 2-2h10a2.003 2.003 0 0 1 2 2v8H5V7zm16 10a2.003 2.003 0 0 1-2 2H5a2.003 2.003 0 0 1-2-2v-1h18v1z"/>`),
		g.Group(children),
	)
}
package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsanaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.5a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5Zm-5 9a3.75 3.75 0 1 0 0 7.498A3.75 3.75 0 0 0 7 12.5Zm10 0a3.75 3.75 0 1 0 0 7.499a3.75 3.75 0 0 0 0-7.499Z"/>`),
		g.Group(children),
	)
}
package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5zm-10 19H3v-5.5h8.5V21zm0-6.5H3v-5h8.5v5zM21 21h-8.5v-5.5H21V21zm0-6.5h-8.5v-5H21v5zm0-6H3V3h18v5.5z"/>`),
		g.Group(children),
	)
}
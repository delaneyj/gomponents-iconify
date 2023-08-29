package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5zm-13 19H3v-5.5h5.5V21zm0-6.5H3v-5h5.5v5zm0-6H3V3h5.5v5.5zm6 12.5h-5v-5.5h5V21zm0-6.5h-5v-5h5v5zm0-6h-5V3h5v5.5zM21 21h-5.5v-5.5H21V21zm0-6.5h-5.5v-5H21v5zm0-6h-5.5V3H21v5.5z"/>`),
		g.Group(children),
	)
}
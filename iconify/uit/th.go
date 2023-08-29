package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Th(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5zm-10 19H3v-8.5h8.5V21zm0-9.5H3V3h8.5v8.5zM21 21h-8.5v-8.5H21V21zm0-9.5h-8.5V3H21v8.5z"/>`),
		g.Group(children),
	)
}
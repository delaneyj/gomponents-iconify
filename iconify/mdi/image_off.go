package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 17.2L6.8 3H19c1.1 0 2 .9 2 2v12.2m-.3 4.8l-1-1H5c-1.1 0-2-.9-2-2V4.3l-1-1L3.3 2L22 20.7L20.7 22m-3.9-4l-3.9-3.9l-1.9 2.4l-2.5-3L5 18h11.8Z"/>`),
		g.Group(children),
	)
}
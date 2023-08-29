package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 19h-4L5 5v14H2V1h5l8 13c-.02-.84 0-1 0-2V1h3v18Z"/>`),
		g.Group(children),
	)
}
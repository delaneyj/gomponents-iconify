package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtGalleryEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.21 3L5.85.65a.5.5 0 0 0-.707-.003L5.14.65L2.79 3H1.5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5H8.21zM5.5 1.71L6.79 3H4.21L5.5 1.71zM9 9H2V4h7v5zM4.5 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zM8 8H4l.75-1.5l.5 1L6.5 5L8 8z" fill="currentColor"/>`),
		g.Group(children),
	)
}
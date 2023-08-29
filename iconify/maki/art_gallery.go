package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10.71 3L7.85.15a.5.5 0 0 0-.707-.003L7.14.15L4.29 3H1.5a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5zM7.5 1.21L9.29 3H5.71zM13 12H2V4h11zM5 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm7 4H4.5L6 8l1.25 2.5L9.5 6z"/>`),
		g.Group(children),
	)
}
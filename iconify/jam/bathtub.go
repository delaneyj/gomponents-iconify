package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bathtub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.487 2.04A3 3 0 0 1 11 5v2H5V5c0-1.039.528-1.955 1.33-2.493A2 2 0 0 0 3 4v6h16a1 1 0 0 1 0 2v1a6.002 6.002 0 0 1-4 5.659V19a1 1 0 0 1-2 0H7a1 1 0 0 1-2 0v-.341A6.002 6.002 0 0 1 1 13v-1a1 1 0 0 1 0-2V4a4 4 0 0 1 7.487-1.96zM17 12H3v1a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-1zM8 4a1 1 0 0 0-1 1h2a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}
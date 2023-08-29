package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.157 12a4.573 4.573 0 1 1-9.147 0a4.573 4.573 0 0 1 9.147 0zm10.833 0a4.573 4.573 0 1 1-9.147 0a4.573 4.573 0 0 1 9.147 0z"/>`),
		g.Group(children),
	)
}
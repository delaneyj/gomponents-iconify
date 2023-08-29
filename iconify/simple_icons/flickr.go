package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.334 6.666a5.335 5.335 0 0 0 0 10.668A5.333 5.333 0 0 0 10.666 12a5.333 5.333 0 0 0-5.332-5.334zm13.332 0A5.333 5.333 0 0 0 13.334 12A5.333 5.333 0 1 0 24 12a5.335 5.335 0 0 0-5.334-5.334Z"/>`),
		g.Group(children),
	)
}
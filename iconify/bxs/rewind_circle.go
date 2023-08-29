package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.485 2 12s4.486 10 10 10c5.515 0 10-4.485 10-10S17.515 2 12 2zm5 14l-6-4v4l-6-4l6-4v4l6-4v8z"/>`),
		g.Group(children),
	)
}
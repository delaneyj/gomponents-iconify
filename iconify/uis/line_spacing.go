package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8h11c.6 0 1-.4 1-1s-.4-1-1-1H10c-.6 0-1 .4-1 1s.4 1 1 1zm-4.3 7.3V8.7c.2.2.4.3.6.3c.3 0 .5-.1.7-.2c.4-.4.5-1 .1-1.4l-1.7-2C5.2 5.1 5 5 4.7 5s-.6.1-.8.4l-1.7 2c-.3.4-.3 1 .2 1.4c.4.3.9.3 1.3 0v6.6c-.4-.3-.9-.4-1.3 0s-.5 1-.1 1.4l1.7 2c.1.1.4.2.7.2s.6-.1.8-.4l1.7-2c.4-.4.3-1.1-.1-1.4c-.5-.3-1.1-.3-1.4.1zM21 11H10c-.6 0-1 .4-1 1s.4 1 1 1h11c.6 0 1-.4 1-1s-.4-1-1-1zm0 5H10c-.6 0-1 .4-1 1s.4 1 1 1h11c.6 0 1-.4 1-1s-.4-1-1-1z"/>`),
		g.Group(children),
	)
}
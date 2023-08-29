package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.8 13.4c-.1-.1-.1-.2-.2-.2c-.5-.3-1.1-.2-1.5.2l-4 6c0 .2-.1.4-.1.6c0 .6.4 1 1 1h8c.2 0 .4-.1.6-.2c.5-.3.6-.9.3-1.4l-4.1-6zM19 3H5C3.3 3 2 4.3 2 6v9c0 1.7 1.3 3 3 3h.8c.6 0 1-.4 1-1s-.4-1-1-1H5c-.6 0-1-.4-1-1V6c0-.6.4-1 1-1h14c.6 0 1 .4 1 1v9c0 .6-.4 1-1 1h-.8c-.6 0-1 .4-1 1s.4 1 1 1h.8c1.7 0 3-1.3 3-3V6c0-1.7-1.3-3-3-3z"/>`),
		g.Group(children),
	)
}
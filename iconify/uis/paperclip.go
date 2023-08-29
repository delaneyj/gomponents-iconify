package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.1 12.4l-6.2 6.2c-1.7 1.7-4.4 1.7-6 0c-1.7-1.7-1.7-4.4 0-6l8-8c1-.9 2.5-.9 3.5 0c1 1 1 2.6 0 3.5L10.5 15c-.3.3-.8.3-1.1 0c-.3-.3-.3-.8 0-1.1l5.1-5.1c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L8 12.6c-1.1 1.1-1.1 2.8 0 3.9c1.1 1 2.8 1 3.9 0l6.9-6.9c1.8-1.8 1.8-4.6 0-6.4c-1.8-1.8-4.6-1.8-6.4 0l-8 8c-1.2 1.2-1.8 2.8-1.8 4.4c0 3.5 2.8 6.2 6.3 6.2c1.7 0 3.2-.7 4.4-1.8l6.2-6.2c.4-.4.4-1 0-1.4s-1-.4-1.4 0z"/>`),
		g.Group(children),
	)
}
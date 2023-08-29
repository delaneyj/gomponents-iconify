package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserMd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.9 13.2c-.1 0-.1-.1-.2-.1C17.2 12 18 10.3 18 8.4v-.7l.3-2.4c.2-1.6-.9-3-2.4-3.3l-.9-.2c-2-.3-4-.3-6 0l-.8.2c-1.6.3-2.7 1.7-2.5 3.3L6 7.7v.7c0 1.8.8 3.6 2.3 4.7c-.1 0-.1.1-.2.1c-3.3 1.4-5.6 4.5-6 8.1c-.1.5.3 1 .9 1.1c.6.1 17.5 0 18 0h.1c.5-.1.9-.6.9-1.1c-.5-3.6-2.8-6.7-6.1-8.1zM12 16.3l-1.7-1.7c1.1-.2 2.2-.2 3.3 0L12 16.3zm0-3.9c-2.2 0-3.9-1.7-4-3.9h8c-.1 2.2-1.8 3.9-4 3.9z"/>`),
		g.Group(children),
	)
}
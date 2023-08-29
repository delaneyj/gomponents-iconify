package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserNurse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.2 12.9s-.1 0 0 0c-.2-.1-.3-.2-.5-.2c2.6-2 3.1-5.8 1-8.4s-5.8-3.1-8.4-1s-3.1 5.8-1 8.4c.3.4.6.7 1 1c-.1.1-.3.1-.4.2h-.1c-3.2 1.5-5.4 4.5-5.8 8c0 .5.4 1 1 1.1h18c.5-.1.9-.6.9-1.1c-.3-3.5-2.5-6.5-5.7-8zM8 7.6c.2-2.2 2.2-3.8 4.3-3.6c1.9.2 3.4 1.7 3.6 3.6H8zm4 8.6l-1.9-1.9c1.3-.3 2.6-.3 3.9 0l-2 1.9z"/>`),
		g.Group(children),
	)
}
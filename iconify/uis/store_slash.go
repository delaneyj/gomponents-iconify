package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 10.6c.8.9 1.8 1.4 3 1.4v1.8l2 2v-4.3c1.2-.7 2-2 2-3.4c0-.1 0-.3-.1-.4l-2-5c-.1-.5-.5-.7-.9-.7H6.2l8.8 8.6c-.1.1 0 0 0 0zm7.7 10.7L20 18.6l-2-2l-4.8-4.8l-9.1-9.2l-1.4-1.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2 2l-1.2 3c-.1 0-.1.2-.1.3c0 1.4.8 2.7 2 3.4V21c0 .6.4 1 1 1h14c.4 0 .8-.3.9-.7l1.4 1.4c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4zM10 14c-.6 0-1 .4-1 1v5H6v-8c1.2 0 2.2-.5 3-1.4c.3.3.6.6 1 .8l2.6 2.6H10zm8 6h-3v-3.6l3 3v.6z"/>`),
		g.Group(children),
	)
}
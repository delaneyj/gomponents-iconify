package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Illustration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.9 9.6c-.2-.5-.8-.7-1.3-.5l-2.9 1.4l-2.1-2.1l-2.1-2.1l1.4-2.9c.2-.5 0-1.1-.5-1.3c-.5-.2-1.1 0-1.3.5l-1.5 3.1l-6.4 1c-.4.1-.7.4-.8.8L2 19.1c-.1.3 0 .7.3.9L4 21.7c.2.2.5.3.7.3h.2l11.6-2.4c.4-.1.7-.4.8-.8l1-6.4l3.1-1.5c.5-.2.7-.8.5-1.3zm-6.5 8.2l-9.8 2l3.7-3.7c1.5.7 3.3.1 4-1.4s.1-3.3-1.4-4c-1.1-.5-2.5-.3-3.4.6c-.9.9-1.1 2.3-.6 3.4l-3.7 3.7l2-9.8l5.8-1l2.2 2.2l2.2 2.2l-1 5.8zm-5.8-4.4c0-.3.1-.5.3-.7c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4c-.4.4-1 .4-1.4 0c-.2-.2-.3-.5-.3-.7z"/>`),
		g.Group(children),
	)
}
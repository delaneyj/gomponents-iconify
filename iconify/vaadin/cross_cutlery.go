package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossCutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.9 8.6c.6-.1 1.2-.4 1.6-.9l3.1-3.1c.4-.4.4-1 0-1.4l-.1-.2l-3 3c-.2.2-.6.2-.9 0s-.2-.6 0-.9l2.6-2.6c.2-.2.2-.6 0-.9c-.2-.2-.6-.2-.9 0l-2.6 2.6c-.2.2-.6.2-.9 0c-.2-.2-.2-.6 0-.9l3-3l-.1-.1c-.4-.4-1-.4-1.4 0L8.2 3.5c-.4.4-.7 1-.8 1.6L2.5.3c-.4-.4-1-.3-1.3 0L1 .5C-.4 1.9.1 4.7 2.5 7.1l.8.8c.4.4.9.7 1.5.8c-.5.4-.8.8-.8.8L.6 12.9c-.7.7-.7 1.9 0 2.6s1.9.7 2.6 0L6.5 12c.2-.2.7-.8 1.3-1.5c.3.4.5.6.5.6l4.3 4.3c.7.7 1.9.7 2.6 0s.7-1.9 0-2.6l-4.3-4.2z"/>`),
		g.Group(children),
	)
}
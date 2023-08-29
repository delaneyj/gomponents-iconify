package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m10.3 8.2l-.9.9l.9.9l-1.2 1.2l4.3 4.3c.6.6 1.5.6 2.1 0s.6-1.5 0-2.1l-5.2-5.2zm3.9 6.8c-.4 0-.8-.3-.8-.8c0-.4.3-.8.8-.8s.8.3.8.8s-.3.8-.8.8zM3.6 8l.9-.6L6 5.7l.9.9l.9-.9l-.1-.1c.2-.5.3-1 .3-1.6c0-2.2-1.8-4-4-4c-.6 0-1.1.1-1.6.3l2.9 2.9l-2.1 2.1L.3 2.4C.1 2.9 0 3.4 0 4c0 2.1 1.6 3.7 3.6 4z"/><path fill="currentColor" d="m8 10.8l.9-.8l-.9-.9l5.7-5.7l1.2-.4L16 .8l-.7-.7l-2.3 1l-.5 1.2L6.9 8L6 7.1l-.8.9s.8.6-.1 1.5c-.5.5-1.3-.1-2.8 1.4L.2 13s-.6 1 .6 2.2s2.2.6 2.2.6l2.1-2.1c1.4-1.4.9-2.3 1.3-2.7c.9-.9 1.6-.2 1.6-.2zm-3.1-.4l.7.7l-3.8 3.8l-.7-.7z"/>`),
		g.Group(children),
	)
}
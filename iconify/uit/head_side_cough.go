package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadSideCough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20.4c-.3 0-.6.3-.6.6s.3.6.6.6s.6-.3.6-.6s-.3-.6-.6-.6zM23 9c0-3.8-3.1-6.9-6.9-7c-3.8-.1-7 3-7.1 6.9l-2.1 4.4v.2c0 .3.2.5.5.5H9v2c0 1.1.9 2 2 2h1v2.5c0 .3.2.5.5.5s.5-.2.5-.5V18h.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H11c-.6 0-1-.4-1-1v-2.5c0-.3-.2-.5-.5-.5H8.2L10 9.2V9c0-3.3 2.6-6 5.9-6c3.3 0 6 2.6 6.1 5.9l-2 7.4v.2l1 4c.1.2.3.4.5.4h.1c.3-.1.4-.3.4-.6l-1-3.9L23 9c0 .1 0 0 0 0zM2 17.4c-.3 0-.6.3-.6.6s.3.6.6.6s.6-.3.6-.6s-.3-.6-.6-.6zm4-1c-.3 0-.6.3-.6.6s.3.6.6.6s.6-.3.6-.6s-.3-.6-.6-.6z"/>`),
		g.Group(children),
	)
}
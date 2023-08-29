package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.9 11.1c-.4-.2-4-2.3-4-2.3L4.6 1.7c-.2-.2-.6-.2-.9-.1c0 0-.1 0-.1.1c-.4.1-.6.5-.6.9v19c0 .3.2.7.4.8h.1c.1 0 .2.1.3.1c.2 0 .4-.1.6-.2c.4-.2 12.4-7.2 12.4-7.2l4-2.3c.4-.2.6-.5.6-.9c.1-.3-.1-.7-.5-.8zm-5.1-1.8l-2 2l-7.4-7.5l9.4 5.5zM4 21V2.9l9.2 9.1L4 21zm2.4-1l7.4-7.4l2 1.9c-1.6 1.1-6.4 3.8-9.4 5.5zm14-8l-3.7 2.1l-2.1-2.1l2.2-2.1c.8.4 3 1.7 3.6 2.1z"/>`),
		g.Group(children),
	)
}
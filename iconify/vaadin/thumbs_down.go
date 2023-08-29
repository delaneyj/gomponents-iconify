package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.6 7.8s.5.5.4 1.6c0 1.5-1.6 1.6-1.6 1.6H12c-.2 0-.3.2-.3.4c.3.7.8 2.1.6 3.1c-.3 1.4-1.5 1.5-1.9 1.5c-.1 0-.2-.1-.2-.2l-1-2.8s0-.1-.1-.1l-2.6-2.8c-.1-.1-.2-.1-.3-.1H6V3h.2c.7 0 3.2-2 5.4-2s2.7.3 3.1 1c.4.7.1 1.3.1 1.3s.5.3.6 1c.1.7-.1 1.1-.1 1.1s.5.4.5 1.2c.1.9-.2 1.2-.2 1.2zM0 11h5V3H0v8zm2.5-3.5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`),
		g.Group(children),
	)
}
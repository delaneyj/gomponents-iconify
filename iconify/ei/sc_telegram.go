package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScTelegram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M37.1 13L9.4 24c-.9.3-.8 1.6.1 1.9l7 2.2l2.8 8.8c.2.7 1.1.9 1.6.4l4.1-3.8l7.8 5.7c.6.4 1.4.1 1.6-.6l5.4-23.2c.3-1.7-1.2-3-2.7-2.4zM20.9 29.8L20 35l-2-7.2L37.5 15L20.9 29.8z"/>`),
		g.Group(children),
	)
}
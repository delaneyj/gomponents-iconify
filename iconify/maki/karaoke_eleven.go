package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KaraokeEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4.5 4.4l-2 2l-1 1c-.5.4-.6 1.1-.3 1.7l.6.7c.6.4 1.3.2 1.7-.3l1-1l2-2l-2-2.1zM3.1 8.5l-.6-.7l2-2l.7.7l-2.1 2zM5 2l1-1h3l1 1v3L9 6H8L5 3V2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
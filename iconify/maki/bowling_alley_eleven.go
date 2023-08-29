package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlingAlleyEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.02 3.94c.04-.1.13-.4.17-.49A2.825 2.825 0 0 0 4.5 3a3.5 3.5 0 0 0 0 7a2.88 2.88 0 0 0 1.54-.4a8.597 8.597 0 0 1-.4-2.61a12.759 12.759 0 0 1 .38-3.05zM2.5 6.98a.48.48 0 1 1 .48-.48a.487.487 0 0 1-.48.48zm1-2a.48.48 0 1 1 .48-.48a.487.487 0 0 1-.48.48zm1 2a.48.48 0 1 1 .48-.48a.487.487 0 0 1-.48.48zM8.098 2.5c0-.5-.1-.5-.1-1c0-.782.5-.7.5-.7s.5-.082.5.7c0 .5-.1.5-.1 1c0 1.5 1.075 2.503 1.075 4.5c0 1.5-.475 3-.975 3h-1c-.5 0-.975-1.5-.975-3c0-1.997 1.075-3 1.075-4.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}
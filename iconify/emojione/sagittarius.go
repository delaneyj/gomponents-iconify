package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sagittarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#428bc1"/><path fill="#fff" stroke="#fff" stroke-miterlimit="10" d="M48 17.4c-.1-.8-.6-1.4-1.4-1.4H32c-1.9 0-1.9 2.9 0 2.9h11L24.7 37.2L18.5 31c-1.3-1.3-3.4.7-2.1 2.1l6.2 6.2l-6.2 6.2c-.4.5-.5 1.1-.2 1.5c.2.5.6.8 1.2.9h.4c.1 0 .1 0 .2-.1c.1 0 .1-.1.2-.1c.1-.1.2-.1.2-.2l6.2-6.2l6.2 6.2c1.3 1.3 3.4-.7 2.1-2.1l-6.2-6.2c6.2-6 12.3-12.1 18.4-18.2v11c0 1.9 2.9 1.9 2.9 0V17.4z"/>`),
		g.Group(children),
	)
}
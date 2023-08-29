package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomensRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ff5a79"/><path fill="#fff" d="M36 23h-8c-1 0-2 1-2 2l-3 11h3l2-9v4l-2 9h2v11h3V40h2v11h3V40h2l-2-9v-4l2 9h3l-3-11c0-1-1-2-2-2m0-4c0 1-1 2-2 2h-4c-1 0-2-1-2-2v-4c0-1 1-2 2-2h4c1 0 2 1 2 2v4"/>`),
		g.Group(children),
	)
}
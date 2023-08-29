package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 10h-.06a9 9 0 0 0-17.88 0H7a5 5 0 0 0 0 10h2v-9a7 7 0 0 1 14 0v10a4 4 0 0 1-3.17 3.91a4 4 0 1 0 .05 2A6 6 0 0 0 25 21v-1a5 5 0 0 0 0-10ZM4 15a3 3 0 0 1 3-3v6a3 3 0 0 1-3-3Zm12 13a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm9-10v-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}
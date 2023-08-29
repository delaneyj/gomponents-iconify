package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 5H8L4 19v8h24v-8L24 5zm3 21H5v-7h8.75a2.25 2.25 0 0 0 4.5 0H27v7z"/>`),
		g.Group(children),
	)
}
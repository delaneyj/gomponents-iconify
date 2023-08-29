package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yoast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20.8 4l-5 13.5L13 10h-3l4 9.5s.2.6.3 1.1v.7c0 .2-.1.5-.2.7c-.6 1.2-1.7 1.9-3 1.9v3c2.7 0 5.1-1.7 6-4.2L24 4h-3.2zM9 7c-2.2 0-4 1.8-4 4v10c0 2.2 1.8 4 4 4h.9v-1H9c-1.7 0-3-1.4-3-3V11c0-1.7 1.3-3 3-3h9.1l.4-1H9zm15 .1l-.3 1C25 8.4 26 9.6 26 11v13h-8.3c-.2.3-.4.7-.6 1H27V11c0-1.9-1.3-3.4-3-3.9z"/>`),
		g.Group(children),
	)
}
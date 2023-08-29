package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMicronesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#428bc1"/><path fill="#fff" d="m32.8 21.6l2.4 1.8l-.9-2.9l2.4-1.9h-3l-.9-2.9l-.9 2.9h-3l2.4 1.9l-.9 2.9zm0 20.8l2.4-1.8l-.9 2.9l2.4 1.9h-3l-.9 2.9l-.9-2.9h-3l2.4-1.9l-.9-2.9zM21.9 29.9l2.4-1.8l-.9 3l2.4 1.8h-3l-.9 3l-.9-3h-3l2.4-1.8l-.9-3zm20.2 0l-2.4-1.8l.9 3l-2.4 1.8h3l.9 3l.9-3h3l-2.4-1.8l.9-3z"/>`),
		g.Group(children),
	)
}
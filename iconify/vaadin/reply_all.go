package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8c0-5-4.9-5-4.9-5H9V0L3 6l6 6V9h2.2c3.5 0 1.8 7 1.8 7s3-4.1 3-8z"/><path fill="currentColor" d="m0 6l6 6v-1.5L1.5 6L6 1.5V0z"/>`),
		g.Group(children),
	)
}
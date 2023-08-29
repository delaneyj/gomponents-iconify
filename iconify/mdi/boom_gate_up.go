package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoomGateUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.9 3c-.6-1-1.8-1.3-2.7-.7L6.9 8.2C6.6 8.1 6.3 8 6 8c-1.7 0-3 1.3-3 3v9c-.6 0-1 .4-1 1v1h8v-1c0-.6-.4-1-1-1v-8.4l10.1-5.8c1-.6 1.3-1.8.8-2.8M6 12.5c-.8 0-1.5-.7-1.5-1.5S5.2 9.5 6 9.5s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5m4.4-2.9l-2.3-1l1.3-.8l2.3 1l-1.3.8m3.5-2l-2.3-1l1.3-.8l2.3 1l-1.3.8m3.4-2l-2.3-1l1.3-.8l2.3 1l-1.3.8"/>`),
		g.Group(children),
	)
}
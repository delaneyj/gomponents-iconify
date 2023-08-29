package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 132v76a20 20 0 0 1-20 20H48a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h76a12 12 0 0 1 0 24H52v144h144v-72a12 12 0 0 1 24 0Zm16-72a40 40 0 1 1-40-40a40 40 0 0 1 40 40Zm-24 0a16 16 0 1 0-16 16a16 16 0 0 0 16-16Z"/>`),
		g.Group(children),
	)
}
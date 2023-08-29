package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactlessPaymentBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152.58 66.35a130.61 130.61 0 0 1 0 123.3a12 12 0 1 1-21.17-11.3a106.7 106.7 0 0 0 0-100.7a12 12 0 1 1 21.16-11.3Zm-52.22 11.06a12 12 0 0 0-5 16.23a73 73 0 0 1 0 68.72a12 12 0 1 0 21.18 11.28a97 97 0 0 0 0-91.28a12 12 0 0 0-16.18-4.95ZM236 128A108 108 0 1 1 128 20a108.12 108.12 0 0 1 108 108Zm-24 0a84 84 0 1 0-84 84a84.09 84.09 0 0 0 84-84Z"/>`),
		g.Group(children),
	)
}
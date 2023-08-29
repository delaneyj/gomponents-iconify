package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21L21 3M5.831 14.161A15.946 15.946 0 0 1 3 6a2 2 0 0 1 2-2h4l2 5l-2.5 1.5c.108.22.223.435.345.645m1.751 2.277A11.03 11.03 0 0 0 13.5 15.5L15 13l5 2v4a2 2 0 0 1-2 2a15.963 15.963 0 0 1-10.344-4.657"/>`),
		g.Group(children),
	)
}
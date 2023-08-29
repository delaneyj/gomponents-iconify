package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><path fill="#1565C0" d="m21 23l-7 10h14z"/><path fill="#1976D2" d="M28 26.4L23 33h10z"/><circle cx="31.5" cy="24.5" r="1.5" fill="#1976D2"/>`),
		g.Group(children),
	)
}
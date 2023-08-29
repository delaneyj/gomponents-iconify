package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><path fill="#1976D2" d="M16 21h17v2H16zm0 4h13v2H16zm0 4h17v2H16zm0 4h13v2H16z"/>`),
		g.Group(children),
	)
}
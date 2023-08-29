package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGift0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" stroke-linecap="round" d="M41 44V20H7v24h34Z"/><path stroke-linecap="round" d="M24 44V20m17 24H7"/><path fill="#555" d="M4 12h40v8H4z"/><path stroke-linecap="round" d="m16 4l8 8l8-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGift0)"/>`),
		g.Group(children),
	)
}
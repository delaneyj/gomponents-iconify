package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGift0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linecap="round" d="M41 44V20H7v24h34Z"/><path stroke="#000" stroke-linecap="round" d="M24 44V20"/><path stroke="#fff" stroke-linecap="round" d="M41 44H7"/><path fill="#fff" stroke="#fff" d="M4 12h40v8H4z"/><path stroke="#fff" stroke-linecap="round" d="m16 4l8 8l8-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGift0)"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 24h12v20H29zM7 24h14v20H7zm3-12.546V24h8V4c-6.5 0-8 5.636-8 7.454ZM7 32h14"/>`),
		g.Group(children),
	)
}
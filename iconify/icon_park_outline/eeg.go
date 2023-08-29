package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 16V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v7m0 16v7a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-7M5 24h8.075L20 16l7 16l6.975-8H43"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swappa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.93 32.043A21.457 21.457 0 0 0 27 2.733M7.079 37.234a21.452 21.452 0 0 0 33.842 0M21 2.733a21.457 21.457 0 0 0-16.93 29.31"/><circle cx="17.5" cy="24" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.5" cy="24" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
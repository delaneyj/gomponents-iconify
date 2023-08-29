package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 9.731L4.5 24h28.632L43.5 9.731zm-7.497 2.742l1.18-4.736L4.5 24m39 14.269L4.5 24h28.632L43.5 38.269z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.003 35.526l1.18 4.737L4.5 24"/>`),
		g.Group(children),
	)
}
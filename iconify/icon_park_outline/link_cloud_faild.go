package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkCloudFaild(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 33c-3.333 0-8-1.5-8-7.5c0-7 7-8.5 9-8.5c1-3.5 3-9 11-9c7 0 10 4 11 7.5c0 0 9 1 9 9.5c0 6-4 8-8 8m-7-5L19 38m0-10l10 10"/>`),
		g.Group(children),
	)
}
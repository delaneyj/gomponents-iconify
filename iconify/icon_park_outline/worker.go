package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 16a8 8 0 1 1-16 0m8-8a8 8 0 0 0-8 8h16a8 8 0 0 0-8-8Zm-12 8h24M24 4v4m0 19c-9.389 0-17 7.163-17 16h34c0-8.837-7.611-16-17-16Zm-6 7v4m12-4v4"/>`),
		g.Group(children),
	)
}
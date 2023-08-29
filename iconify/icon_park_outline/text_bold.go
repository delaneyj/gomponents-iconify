package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 24c5.506 0 9.969-4.477 9.969-10S29.506 4 24 4H11v20h13Zm4.031 20C33.537 44 38 39.523 38 34s-4.463-10-9.969-10H11v20h17.031Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
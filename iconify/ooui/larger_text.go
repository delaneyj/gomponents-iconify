package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargerText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.66 18h-2a.85.85 0 0 1-.56-.17a1.11 1.11 0 0 1-.32-.43l-1.33-3.53h-6.9L5.22 17.4a1.06 1.06 0 0 1-.31.41a.83.83 0 0 1-.56.19h-2L8.68 2h2.63zm-4.92-6l-2.2-5.84A16.17 16.17 0 0 1 10 4.43q-.12.52-.27 1t-.27.77L7.26 12z"/>`),
		g.Group(children),
	)
}
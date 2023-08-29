package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.71 20.29l-16-16l-2-2a1 1 0 0 0-1.42 1.42L4 5.41v6.47a9 9 0 0 0 3.77 7.32l3.65 2.61a1 1 0 0 0 1.16 0l3.65-2.61a8.21 8.21 0 0 0 .86-.7l3.2 3.21a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-6.64-2.71L12 19.77l-3.07-2.19A7 7 0 0 1 6 11.88V7.41l9.67 9.68c-.19.17-.39.33-.6.49Zm-5-12.51A10.15 10.15 0 0 0 12 4.19a9.82 9.82 0 0 0 6 1.39v6.3a6.88 6.88 0 0 1-.1 1.18a1 1 0 0 0 .83 1.15h.16a1 1 0 0 0 1-.84a9.77 9.77 0 0 0 .12-1.5V4.43a1 1 0 0 0-.37-.77a1 1 0 0 0-.83-.21a7.89 7.89 0 0 1-6.23-1.27a1 1 0 0 0-1.14 0a8 8 0 0 1-2 1a1 1 0 1 0 .64 1.9Z"/>`),
		g.Group(children),
	)
}
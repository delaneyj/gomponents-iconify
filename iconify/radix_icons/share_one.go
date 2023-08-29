package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 7.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm.713 1.164a2.5 2.5 0 1 1 0-2.328l3.391-2.12A2.5 2.5 0 1 1 14 3.5a2.5 2.5 0 0 1-4.484 1.52l-3.53 2.207a2.526 2.526 0 0 1 0 .546l3.53 2.206a2.5 2.5 0 1 1-.411.804l-3.392-2.12ZM11.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
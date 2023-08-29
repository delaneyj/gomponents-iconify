package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropdownMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.1a.4.4 0 1 0 0 .8h7a.4.4 0 0 0 0-.8h-7Zm0 2a.4.4 0 1 0 0 .8h7a.4.4 0 0 0 0-.8h-7Zm-.4 2.4c0-.22.18-.4.4-.4h7a.4.4 0 0 1 0 .8h-7a.4.4 0 0 1-.4-.4Zm.4 1.6a.4.4 0 1 0 0 .8h7a.4.4 0 0 0 0-.8h-7Zm-.4 2.4c0-.22.18-.4.4-.4h7a.4.4 0 0 1 0 .8h-7a.4.4 0 0 1-.4-.4ZM2.5 9.25L5 6H0l2.5 3.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
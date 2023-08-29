package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tokens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM3 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0ZM10.5 2a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM9 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm-7 6a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM4.5 9a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm6-1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM9 10.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
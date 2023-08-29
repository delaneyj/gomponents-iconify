package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm-2.829 7A4.491 4.491 0 0 1 6 7.5c0-1.414.652-2.675 1.671-3.5H4.5a3.5 3.5 0 1 0 0 7h3.171ZM0 7.5A4.5 4.5 0 0 1 4.5 3h6a4.5 4.5 0 1 1 0 9h-6A4.5 4.5 0 0 1 0 7.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
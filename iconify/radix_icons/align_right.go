package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 1a.5.5 0 0 0-.5.5V6H3a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h11v4.5a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
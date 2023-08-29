package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHandleVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.1 12.5a.4.4 0 0 0 .8 0v-10a.4.4 0 0 0-.8 0v10Zm2 0a.4.4 0 0 0 .8 0v-10a.4.4 0 0 0-.8 0v10Zm2.4.4a.4.4 0 0 1-.4-.4v-10a.4.4 0 1 1 .8 0v10a.4.4 0 0 1-.4.4Zm1.6-.4a.4.4 0 0 0 .8 0v-10a.4.4 0 0 0-.8 0v10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
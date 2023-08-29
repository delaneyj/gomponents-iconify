package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .765.424L7.5 11.59l3.735 2.334A.5.5 0 0 0 12 13.5v-11a.5.5 0 0 0-.5-.5h-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
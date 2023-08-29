package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 3h13v1H1V3Zm0 3h13v2H1V6Zm13 4.25H1v2.5h13v-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
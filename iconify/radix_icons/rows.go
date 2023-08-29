package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 12.85H1v1.3h13v-1.3Zm0-4H1v1.3h13v-1.3Zm-13-4h13v1.3H1v-1.3Zm13-4H1v1.3h13V.85Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
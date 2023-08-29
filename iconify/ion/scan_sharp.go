package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M388 466h-68v-44h68a34 34 0 0 0 34-34v-68h44v68a78.09 78.09 0 0 1-78 78Zm78-274h-44v-68a34 34 0 0 0-34-34h-68V46h68a78.09 78.09 0 0 1 78 78ZM192 466h-68a78.09 78.09 0 0 1-78-78v-68h44v68a34 34 0 0 0 34 34h68ZM90 192H46v-68a78.09 78.09 0 0 1 78-78h68v44h-68a34 34 0 0 0-34 34Z"/>`),
		g.Group(children),
	)
}
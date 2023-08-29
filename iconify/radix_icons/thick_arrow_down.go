package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThickArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V6h2.5a.5.5 0 0 1 .407.79l-5 7a.5.5 0 0 1-.814 0l-5-7A.5.5 0 0 1 2.5 6H5V3.5ZM6 4v2.5a.5.5 0 0 1-.5.5H3.472L7.5 12.64L11.528 7H9.5a.5.5 0 0 1-.5-.5V4H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
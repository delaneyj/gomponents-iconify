package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThickArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1a.5.5 0 0 1 .407.21l5 7A.5.5 0 0 1 12.5 9H10v2.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V9H2.5a.5.5 0 0 1-.407-.79l5-7A.5.5 0 0 1 7.5 1ZM3.472 8H5.5a.5.5 0 0 1 .5.5V11h3V8.5a.5.5 0 0 1 .5-.5h2.028L7.5 2.36L3.472 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
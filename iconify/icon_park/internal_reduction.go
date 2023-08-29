package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternalReduction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 42H40C41.1046 42 42 41.1046 42 40V8C42 6.89543 41.1046 6 40 6H8C6.89543 6 6 6.89543 6 8V40C6 41.1046 6.89543 42 8 42Z"/><path fill="#2F88FF" fill-rule="evenodd" d="M42 8C42 6.89543 41.1046 6 40 6H28V20H42V8Z" clip-rule="evenodd"/><path d="M23 25L13 35M13 35V28M13 35H20"/></g>`),
		g.Group(children),
	)
}
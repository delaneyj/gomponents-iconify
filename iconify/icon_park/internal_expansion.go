package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternalExpansion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 42H40C41.1046 42 42 41.1046 42 40V8C42 6.89543 41.1046 6 40 6H8C6.89543 6 6 6.89543 6 8V40C6 41.1046 6.89543 42 8 42Z"/><path fill="#2F88FF" fill-rule="evenodd" d="M42 8C42 6.89543 41.1046 6 40 6H28V20H42V8Z" clip-rule="evenodd"/><path d="M13 35L23 25M23 25V32M23 25H16"/></g>`),
		g.Group(children),
	)
}
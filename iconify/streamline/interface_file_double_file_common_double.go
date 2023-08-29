package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFileDoubleFileCommonDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V1.5a1 1 0 0 1 1-1h4.5l3 3Z"/><path d="M9.5 13.5h-7a1 1 0 0 1-1-1v-9"/></g>`),
		g.Group(children),
	)
}
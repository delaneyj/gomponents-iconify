package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapSingle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.05 3.114c2.143 0 4.09.843 5.526 2.216L16.16 6.744a5.98 5.98 0 0 0-4.112-1.63a5.98 5.98 0 0 0-4.21 1.725L6.424 5.425a7.974 7.974 0 0 1 5.625-2.311Zm-1.073 8.772a1 1 0 1 1 2 0v2h-2v-2Z"/><path fill-rule="evenodd" d="M11.977 6.886a5 5 0 0 0-5 5v4a5 5 0 0 0 10 0v-4a5 5 0 0 0-5-5Zm3 9v-4a3 3 0 0 0-6 0v4a3 3 0 0 0 6 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
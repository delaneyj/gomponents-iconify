package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarHorizontalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 104v48H40v-48Z" opacity=".2"/><path d="M216 96h-40V56a8 8 0 0 0-8-8H48v-8a8 8 0 0 0-16 0v176a8 8 0 0 0 16 0v-8h88a8 8 0 0 0 8-8v-40h72a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8Zm-56-32v32H48V64Zm-32 128H48v-32h80Zm80-48H48v-32h160Z"/></g>`),
		g.Group(children),
	)
}
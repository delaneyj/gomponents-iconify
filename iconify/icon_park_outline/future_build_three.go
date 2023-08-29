package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FutureBuildThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="m20 8l4-4l4 4v36h-8V8Z"/><path stroke-linecap="round" d="m12 20l8-8v32h-8V20ZM4 35l8-7v16H4v-9Zm24-23l8 8v24h-8V12Zm8 16l8 6.5V44h-8V28Z"/></g>`),
		g.Group(children),
	)
}
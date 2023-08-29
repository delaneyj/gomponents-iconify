package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FutureBuildThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M20 8L24 4L28 8V44H20V8Z"/><path stroke-linecap="round" d="M12 20L20 12V44H12V20Z"/><path stroke-linecap="round" d="M4 35L12 28V44H4V35Z"/><path stroke-linecap="round" d="M28 12L36 20V44H28V12Z"/><path stroke-linecap="round" d="M36 28L44 34.5V44H36V28Z"/></g>`),
		g.Group(children),
	)
}
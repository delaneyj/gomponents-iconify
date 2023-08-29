package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M16 6H8C6.89543 6 6 6.89543 6 8V16"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 42H8C6.89543 42 6 41.1046 6 40V32"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 42H40C41.1046 42 42 41.1046 42 40V32"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 6H40C41.1046 6 42 6.89543 42 8V16"/><rect width="20" height="20" x="14" y="14" fill="#2F88FF" rx="2"/></g>`),
		g.Group(children),
	)
}
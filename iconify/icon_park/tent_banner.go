package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentBanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M7 43H39L23 17L7 43Z"/><path d="M18.0769 25L23 28L27.9231 25"/><path d="M23 17V12V4"/><path fill="#2F88FF" d="M35 4H23V12H35L32 8L35 4Z"/></g>`),
		g.Group(children),
	)
}
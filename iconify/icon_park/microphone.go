package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="16" height="28" x="16" y="4" stroke-linecap="round" stroke-linejoin="round" rx="8"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 21V24C10 31.732 16.268 38 24 38V38C31.732 38 38 31.732 38 24V21"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 5V11"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 16H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M27 16H32"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 22H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M27 22H32"/><path d="M24 38V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 44H32"/></g>`),
		g.Group(children),
	)
}
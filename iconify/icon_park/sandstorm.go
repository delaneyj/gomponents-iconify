package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandstorm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 22H36C40.4183 22 44 18.4183 44 14C44 9.58172 40.4183 6 36 6C31.5817 6 28 9.58172 28 14"/><path d="M10 29H4"/><path d="M27 29H21"/><path d="M44 29H38"/><path d="M10 35H4"/><path d="M27 35H21"/><path d="M44 35H38"/><path d="M16 42H4"/><path d="M44 42H32"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="24" cy="30" r="14" fill="#2F88FF" stroke="#000"/><path stroke="#000" stroke-linejoin="round" d="M31 18V11C31 7.13401 27.866 4 24 4V4C20.134 4 17 7.13401 17 11V18"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24 26L24 34"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WavesRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M24 11V37"/><path d="M33 4V44"/><path d="M6 11V37"/><path d="M42 15V33"/><path d="M15 18V30"/></g>`),
		g.Group(children),
	)
}
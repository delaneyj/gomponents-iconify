package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WavesLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M24 17V31"/><path d="M33 11V37"/><path d="M6 17V31"/><path d="M42 18V30"/><path d="M15 4V44"/></g>`),
		g.Group(children),
	)
}
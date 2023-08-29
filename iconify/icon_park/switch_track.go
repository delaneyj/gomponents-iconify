package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchTrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 4V18L8 30L8 44"/><path d="M8 22V17"/><path d="M8 9V4"/><path d="M40 44V39"/><path d="M40 31V26"/></g>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storygraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="7.603" height="29.842" x="5.5" y="6.617" rx="1" ry="1"/><rect width="7.603" height="29.842" x="16.564" y="6.617" rx="1" ry="1"/><rect width="7.603" height="29.842" x="30.911" y="6.343" rx="1" ry="1" transform="rotate(-16.906 34.712 21.264)"/><rect width="37" height="2.828" x="5.5" y="39.043" rx="1" ry="1"/></g>`),
		g.Group(children),
	)
}
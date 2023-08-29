package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherSnowFlakeWinterFreezeSnowFreezingIceColdWeatherSnowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 .5l2 2l2-2M.5 9l2-2l-2-2M9 13.5l-2-2l-2 2M13.5 5l-2 2l2 2m-10-5.5L5 5m0 4l-1.5 1.5m7-7L9 5m0 4l1.5 1.5M7 2.5v9M2.5 7h9"/>`),
		g.Group(children),
	)
}
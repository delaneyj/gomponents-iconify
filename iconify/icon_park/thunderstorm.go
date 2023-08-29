package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thunderstorm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9.45455 30.9942C6.14242 28.461 4 24.4278 4 19.8851C4 12.2166 10.1052 6 17.6364 6C23.9334 6 29.2336 10.3462 30.8015 16.2533C32.0353 15.6159 33.431 15.2567 34.9091 15.2567C39.9299 15.2567 44 19.4011 44 24.5135C44 28.3094 41.7562 31.5716 38.5455 33"/><path d="M17.4141 22.5859L14.5856 25.4144"/><path d="M26.9996 24L19 32.0012H29.004L21.0003 40.0018"/><path d="M33.4141 38.5859L30.5856 41.4144"/></g>`),
		g.Group(children),
	)
}
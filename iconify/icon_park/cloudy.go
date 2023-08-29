package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M9.45455 30.9942C6.14242 28.461 4 24.4278 4 19.8851C4 12.2166 10.1052 6 17.6364 6C23.9334 6 29.2336 10.3462 30.8015 16.2533C32.0353 15.6159 33.431 15.2567 34.9091 15.2567C39.9299 15.2567 44 19.4011 44 24.5135C44 28.3094 41.7562 31.5716 38.5455 33"/><path fill="#2F88FF" d="M26 42C29.3137 42 32 39.3137 32 36C32 32.6863 29.3137 30 26 30C22.6863 30 20 32.6863 20 36C20 39.3137 22.6863 42 26 42Z"/><path stroke-linecap="round" d="M22.2426 15.7574C21.1569 14.6716 19.6569 14 18 14C14.6863 14 12 16.6863 12 20C12 21.6569 12.6716 23.1569 13.7574 24.2426"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thunderbolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 26C4 24.8954 4.89543 24 6 24H42C43.1046 24 44 24.8954 44 26V31.3604C44 32.3138 43.3167 33.1412 42.4802 33.5984C41.2932 34.2471 39.7561 35.5405 39.2065 38.0136C38.9669 39.0919 38.1046 40 37 40H11C9.89543 40 9.03308 39.0919 8.79348 38.0136C8.24394 35.5405 6.70677 34.2471 5.51983 33.5984C4.68326 33.1412 4 32.3138 4 31.3604V26Z"/><path stroke="#fff" d="M14 31H34"/><path stroke="#000" d="M25 4L19 11H29L23 18"/></g>`),
		g.Group(children),
	)
}
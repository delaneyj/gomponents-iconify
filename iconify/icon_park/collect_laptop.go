package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M21 9H11C9.34315 9 8 10.3431 8 12V33H40V23"/><path fill="#2F88FF" d="M4 33H44V35C44 38.3137 41.3137 41 38 41H10C6.68629 41 4 38.3137 4 35V33Z"/><path fill="#2F88FF" stroke-linecap="round" d="M32.3 7C30.4775 7 29 8.43473 29 10.2046C29 13.4091 32.9 16.3223 35 17C37.1 16.3223 41 13.4091 41 10.2046C41 8.43473 39.5225 7 37.7 7C36.5839 7 35.5972 7.53804 35 8.3616C34.4028 7.53804 33.4161 7 32.3 7Z"/></g>`),
		g.Group(children),
	)
}
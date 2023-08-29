package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TajMahal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M16 28H8C6.89543 28 6 28.8954 6 30V44"/><path stroke="#000" stroke-linejoin="round" d="M32 28H40C41.1046 28 42 28.8954 42 30V44"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M19.9994 24H27.9994C27.9994 24 33.1652 18.3215 31.9994 15C31.2446 12.8494 29.5615 11.6597 27.9994 10C26.4373 8.34027 23.9994 6 23.9994 6C23.9994 6 21.5615 8.34027 19.9994 10C18.4373 11.6597 16.7542 12.8494 15.9994 15C14.8337 18.3215 19.9994 24 19.9994 24Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M4 44H44"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 24H24H32V44H16V24Z"/><path stroke="#fff" stroke-linecap="round" d="M24 34V44"/><path stroke="#000" stroke-linecap="round" d="M24 4V7"/><path stroke="#000" stroke-linecap="round" d="M10 24V28"/><path stroke="#000" stroke-linecap="round" d="M38 24V28"/><path stroke="#000" stroke-linecap="round" d="M20 44L28 44"/></g>`),
		g.Group(children),
	)
}
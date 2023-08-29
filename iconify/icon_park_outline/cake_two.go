package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M40 17.9L31.058 8L9 17.9V30h31V17.9Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 18h17.386c.063 0 .114.051.114.114v3.511a2.875 2.875 0 0 0 2.875 2.875v0a2.875 2.875 0 0 0 2.875-2.875v-3.511c0-.063.051-.114.114-.114H40"/><path d="M9.5 23.957c-.602.3-1.162.62-1.678.956C5.418 26.481 4 28.412 4 30.5c0 5.247 8.954 9.5 20 9.5s20-4.253 20-9.5c0-2.14-1.488-4.113-4-5.701"/></g>`),
		g.Group(children),
	)
}
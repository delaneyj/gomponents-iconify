package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M21 9H11a3 3 0 0 0-3 3v21h32V23"/><path d="M4 33h40v2a6 6 0 0 1-6 6H10a6 6 0 0 1-6-6v-2Z"/><path stroke-linecap="round" d="M32.3 7C30.478 7 29 8.435 29 10.205c0 3.204 3.9 6.117 6 6.795c2.1-.678 6-3.59 6-6.795C41 8.435 39.523 7 37.7 7A3.326 3.326 0 0 0 35 8.362A3.326 3.326 0 0 0 32.3 7Z"/></g>`),
		g.Group(children),
	)
}
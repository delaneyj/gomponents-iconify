package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungCameraExpertRaw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.284 9.824H8.716A4.216 4.216 0 0 0 4.5 14.04v19.92a4.216 4.216 0 0 0 4.216 4.216h30.568a4.216 4.216 0 0 0 4.216-4.217V14.04a4.216 4.216 0 0 0-4.216-4.216Z"/><circle cx="37.703" cy="15.622" r="2.635" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="7.905" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.195 21.93h-2.39L21.61 24l1.195 2.07h2.39L26.39 24l-1.195-2.07zM26.39 24l3.217 5.573M18.393 18.427L21.61 24m4.412-7.642l-3.217 5.572m2.39 4.14l-3.217 5.572m.827-5.572h-6.436m8.826-4.14h6.436"/>`),
		g.Group(children),
	)
}
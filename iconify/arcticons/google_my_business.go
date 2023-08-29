package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleMyBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 19.557H24V7.094h-8.142L14.25 19.557z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.858 7.094H9.142a2 2 0 0 0-1.941 1.52L4.5 19.557h9.75m9.75 0a4.875 4.875 0 0 1-9.75 0m0 0a4.875 4.875 0 0 1-9.75 0m29.25 0H24V7.094h8.142l1.608 12.463zM32.142 7.094h6.716a2 2 0 0 1 1.941 1.52L43.5 19.557h-9.75m-9.75 0a4.875 4.875 0 0 0 9.75 0m0 0a4.875 4.875 0 0 0 9.75 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.325 23.977v14.93a2 2 0 0 0 2 2h29.35a2 2 0 0 0 2-2v-14.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.5 33.569h3.956a3.977 3.977 0 0 1-3.88 4.072l-.077.001a4.073 4.073 0 1 1 0-8.147a4.011 4.011 0 0 1 2.02.536"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceUnlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.511 4.5h-18a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-26m-11-11l11 11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.351 33.354h4.874l-2.512-9.362a4.555 4.555 0 1 0-3.412-.004l-2.503 9.366h4.874M8.794 20.686h10.61m9.104 0H39.34m-30.611 3.1h13.128m4.357 0h12.932M8.489 27.107h12.728m5.406 0h12.61"/>`),
		g.Group(children),
	)
}
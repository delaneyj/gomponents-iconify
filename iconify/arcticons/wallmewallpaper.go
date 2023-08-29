package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallmewallpaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.267 32.059l7.694-3.806l6.405 3.037l-7.694 3.806l-6.405-3.037zm-1.401 6.97l5.45-2.695l4.536 2.151l-5.449 2.695l-4.537-2.151zm8.364 2.515l3.635-1.824l3.027 1.455l-3.636 1.824l-3.026-1.455zM24 24.251v2.066M24 5l9.625 9.626L24 24.251l-9.625-9.625zm4.813 14.439l-9.626-9.626m0 9.626l9.626-9.626"/>`),
		g.Group(children),
	)
}
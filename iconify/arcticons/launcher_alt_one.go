package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LauncherAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="23" height="39" x="12.5" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.806 24.837v4.42h-7.612v-4.42m9.376-1.066L24 18.201l-5.57 5.57M12.5 8h23m0 32h-23"/>`),
		g.Group(children),
	)
}
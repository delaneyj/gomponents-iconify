package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLightingLightBulbOnLightingLightShineIncandescentBulbLights(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8a3.5 3.5 0 1 0-5.06 3.12v1.72a.39.39 0 0 0 .39.38h2.34a.39.39 0 0 0 .39-.38v-1.75A3.5 3.5 0 0 0 10.5 8ZM7 .81v1.5m4-.07L9.94 3.31m3.34 2.03h-1.5M3 2.24l1.06 1.07M.72 5.34h1.5"/>`),
		g.Group(children),
	)
}
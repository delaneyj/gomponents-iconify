package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLightingLightBulbLightingLightIncandescentBulbLights(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 5A4.5 4.5 0 1 0 5 9v1.5a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5V9a4.48 4.48 0 0 0 2.5-4ZM5 13.5h4"/>`),
		g.Group(children),
	)
}
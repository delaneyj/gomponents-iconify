package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLightingBrightnessThreeBrightAdjustBrightnessAdjustmentSunRaiseControlsDotSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r=".5"/><path d="M7 12.5v-1m3.89-.61l-.71-.71M12.5 7h-1m-.61-3.89l-.71.71M7 1.5v1m-3.89.61l.71.71M1.5 7h1m.61 3.89l.71-.71"/></g>`),
		g.Group(children),
	)
}
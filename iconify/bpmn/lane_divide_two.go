package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaneDivideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M96 294v1460h1856V294H96zm96 90.758h1660V974h-26v100h26v582.299H192V1074h36V974h-36V384.758zM328 974v100h200V974H328zm300 0v100h200V974H628zm300 0v100h200V974H928zm300 0v100h200V974h-200zm300 0v100h200V974h-200z"/>`),
		g.Group(children),
	)
}
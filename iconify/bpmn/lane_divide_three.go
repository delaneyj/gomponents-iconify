package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaneDivideThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M96 294v1460h1856V294H96zm96 90.758h1660V737h-26v100h26v374h-26v100h26v345.299H192V1311h36v-100h-36V837h36V737h-36V384.758zM328 737v100h200V737H328zm300 0v100h200V737H628zm300 0v100h200V737H928zm300 0v100h200V737h-200zm300 0v100h200V737h-200zM328 1211v100h200v-100H328zm300 0v100h200v-100H628zm300 0v100h200v-100H928zm300 0v100h200v-100h-200zm300 0v100h200v-100h-200z"/>`),
		g.Group(children),
	)
}
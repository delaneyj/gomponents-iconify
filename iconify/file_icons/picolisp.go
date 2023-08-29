package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picolisp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m196.136 19.308l88.596 149.32l-68.687 118.46h-88.597l68.688-118.46l-45.792-77.646l45.792-71.674zM87.953 492.692l42.357-70.923h136.934l-44.298-77.223H87.008L0 492.692h87.953zm375.093-147.57h-69.093l-68.466-116.597l-44.299 75.732l67.97 117.512H512l-48.954-76.647z"/>`),
		g.Group(children),
	)
}
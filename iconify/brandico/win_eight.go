package brandico

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="m985.875 0l-548.75 80.031v397.688h548.75V0zM398.469 85.688L0 143.813v333.906h398.469V85.688zM0 516.376v338.156l398.469 58.813V516.376H0zm437.125 0v402.688l548.75 80.938V516.377h-548.75z"/>`),
		g.Group(children),
	)
}
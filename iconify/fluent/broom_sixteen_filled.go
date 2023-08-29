package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroomSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.854 1.146a.5.5 0 0 1 0 .708L10.896 5.81a4.778 4.778 0 0 0-.708-.707l3.958-3.958a.5.5 0 0 1 .708 0ZM4.653 6.19l-.391.365l5.195 5.195l.396-.396a3.67 3.67 0 0 0 0-5.207c-1.453-1.453-3.765-1.385-5.2.043Zm-3.36 1.855l2.093-.952l5.52 5.52l-.95 2.094a.5.5 0 0 1-.81.146l-6-6a.5.5 0 0 1 .147-.808Z"/>`),
		g.Group(children),
	)
}
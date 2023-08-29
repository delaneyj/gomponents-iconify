package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmptyFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M5.906-.031A1 1 0 0 0 5.781 0a1 1 0 0 0-.718.656L5 .625v.188A1 1 0 0 0 5 1v24a1 1 0 1 0 2 0v-8.531L23.375 9L7 1.531V1A1 1 0 0 0 5.906-.031zM7 3.719L18.563 9L7 14.281V3.72z"/>`),
		g.Group(children),
	)
}
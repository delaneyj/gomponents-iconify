package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmSubtract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28a11 11 0 1 1 11-11a11 11 0 0 1-11 11Zm0-20a9 9 0 1 0 9 9a9 9 0 0 0-9-9ZM4 7.592l3.582-3.589l1.416 1.413l-3.582 3.589zm19-2.184l1.415-1.413l3.581 3.589l-1.415 1.413z"/><path fill="currentColor" d="M11 16h10v2H11z"/>`),
		g.Group(children),
	)
}
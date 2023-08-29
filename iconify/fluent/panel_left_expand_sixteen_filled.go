package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftExpandSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M8.5 7.5h1.791l-.646-.646a.5.5 0 0 1 .707-.707l1.5 1.5a.5.5 0 0 1 0 .707l-1.5 1.5a.5.5 0 1 1-.707-.707l.646-.647H8.5a.5.5 0 0 1 0-1Z"/><path d="M2 4.999a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v6.002a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5Zm10-1H7v8.002h5a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"/></g>`),
		g.Group(children),
	)
}
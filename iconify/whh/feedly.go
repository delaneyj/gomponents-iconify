package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feedly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M633 559L358 832H250l-67-67q-7-7-7-17t7-17l312-310q7-7 17-7t17 7l104 103q7 7 7 17.5t-7 17.5zm0-414L204 576H56L7 528q-7-7-7-17.5T7 493L495 7q7-7 17-7t17 7l104 104q7 7 7 17t-7 17zM391 938l104-104q7-7 17-7t17 7l104 104q7 7 7 17t-7 17l-52 52H443l-52-52q-7-7-7-17t7-17z"/>`),
		g.Group(children),
	)
}
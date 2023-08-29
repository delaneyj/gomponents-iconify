package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroadcastPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.05 3.05a7 7 0 0 0 0 9.9a.5.5 0 0 1-.707.707a8 8 0 0 1 0-11.314a.5.5 0 0 1 .707.707zm2.122 2.122a4 4 0 0 0 0 5.656a.5.5 0 1 1-.708.708a5 5 0 0 1 0-7.072a.5.5 0 0 1 .708.708zm5.656-.708a.5.5 0 0 1 .708 0a5 5 0 0 1 0 7.072a.5.5 0 1 1-.708-.708a4 4 0 0 0 0-5.656a.5.5 0 0 1 0-.708zm2.122-2.12a.5.5 0 0 1 .707 0a8 8 0 0 1 0 11.313a.5.5 0 0 1-.707-.707a7 7 0 0 0 0-9.9a.5.5 0 0 1 0-.707zM6 8a2 2 0 1 1 2.5 1.937V15.5a.5.5 0 0 1-1 0V9.937A2 2 0 0 1 6 8z"/>`),
		g.Group(children),
	)
}
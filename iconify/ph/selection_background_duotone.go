package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionBackgroundDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 48v112a8 8 0 0 1-8 8h-40V96a8 8 0 0 0-8-8H88V48a8 8 0 0 1 8-8h112a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M160 80H48a16 16 0 0 0-16 16v112a16 16 0 0 0 16 16h112a16 16 0 0 0 16-16V96a16 16 0 0 0-16-16Zm0 128H48V96h112ZM136 40a8 8 0 0 1 8-8h16a8 8 0 0 1 0 16h-16a8 8 0 0 1-8-8Zm88 8v8a8 8 0 0 1-16 0v-8h-8a8 8 0 0 1 0-16h8a16 16 0 0 1 16 16Zm0 48v16a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Zm0 56v8a16 16 0 0 1-16 16h-8a8 8 0 0 1 0-16h8v-8a8 8 0 0 1 16 0ZM80 56v-8a16 16 0 0 1 16-16h8a8 8 0 0 1 0 16h-8v8a8 8 0 0 1-16 0Z"/></g>`),
		g.Group(children),
	)
}
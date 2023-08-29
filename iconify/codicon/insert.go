package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m14 1l1 1v4l-1 1H6L5 6V2l1-1h8Zm0 1H6v4h8V2Zm0 7l1 1v4l-1 1H6l-1-1v-4l1-1h8Zm0 1H6v4h8v-4Z" clip-rule="evenodd"/><path d="m1 6.393l1.614 1.614L1 9.62l.694.693L4 8.007L1.694 5.7L1 6.393Z"/></g>`),
		g.Group(children),
	)
}
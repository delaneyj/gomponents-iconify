package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20.385 2.879a3 3 0 0 0-4.243 0L14.02 5l-.707-.708a1 1 0 1 0-1.414 1.415l5.657 5.656A1 1 0 0 0 18.97 9.95l-.707-.707l2.122-2.122a3 3 0 0 0 0-4.242Z"/><path fill-rule="evenodd" d="M11.93 7.091L4.152 14.87a3.001 3.001 0 0 0-.587 3.415L2 19.85l1.414 1.415l1.565-1.566a3.001 3.001 0 0 0 3.415-.586l7.778-7.778L11.93 7.09Zm1.414 4.243L11.93 9.92l-6.364 6.364a1 1 0 0 0 1.414 1.414l6.364-6.364Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.667 2.946a.5.5 0 0 1 .702.087l1.547 1.982a.5.5 0 0 1-.789.615l-1.546-1.982a.5.5 0 0 1 .086-.702Z"/><path d="M16.826 4.926a.5.5 0 0 0-.701.092l-1.547 2.018a.5.5 0 0 0 .794.608l1.546-2.017a.5.5 0 0 0-.092-.701Z"/><path d="M2.975 9.323a4.5 4.5 0 0 1 4.5-4.5h8.5a.5.5 0 0 1 0 1h-8.5a3.5 3.5 0 0 0-3.5 3.5v1.337a.5.5 0 1 1-1 0V9.323Zm2.354 7.731a.5.5 0 0 1-.702-.087l-1.546-1.982a.5.5 0 0 1 .788-.615l1.547 1.982a.5.5 0 0 1-.087.702Z"/><path d="M3.17 15.074a.5.5 0 0 0 .702-.092l1.546-2.018a.5.5 0 1 0-.793-.608l-1.547 2.017a.5.5 0 0 0 .093.701Z"/><path d="M17.021 10.677a4.5 4.5 0 0 1-4.5 4.5h-8.5a.5.5 0 0 1 0-1h8.5a3.5 3.5 0 0 0 3.5-3.5V9.34a.5.5 0 1 1 1 0v1.337Z"/></g>`),
		g.Group(children),
	)
}
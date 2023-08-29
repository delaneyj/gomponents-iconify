package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilRepeatCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.667 5.946a.5.5 0 0 1 .702.087l1.547 1.982a.5.5 0 0 1-.789.615l-1.546-1.982a.5.5 0 0 1 .086-.702Z"/><path d="M19.826 7.926a.5.5 0 0 0-.701.092l-1.547 2.018a.5.5 0 0 0 .794.608l1.546-2.017a.5.5 0 0 0-.092-.701Z"/><path d="M5.975 12.323a4.5 4.5 0 0 1 4.5-4.5h8.5a.5.5 0 0 1 0 1h-8.5a3.5 3.5 0 0 0-3.5 3.5v1.337a.5.5 0 1 1-1 0v-1.337Zm2.354 7.731a.5.5 0 0 1-.702-.087l-1.546-1.982a.5.5 0 0 1 .788-.615l1.547 1.982a.5.5 0 0 1-.087.702Z"/><path d="M6.17 18.074a.5.5 0 0 0 .702-.093l1.546-2.017a.5.5 0 1 0-.793-.608l-1.547 2.017a.5.5 0 0 0 .093.701Z"/><path d="M20.021 13.677a4.5 4.5 0 0 1-4.5 4.5h-8.5a.5.5 0 0 1 0-1h8.5a3.5 3.5 0 0 0 3.5-3.5V12.34a.5.5 0 1 1 1 0v1.337Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilRepeatCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
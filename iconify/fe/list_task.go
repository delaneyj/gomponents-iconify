package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feListTask0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feListTask1" fill="currentColor"><path id="feListTask2" d="M9 13h10a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm0 4h10a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm6-8h4a1 1 0 0 1 0 2h-4a1 1 0 0 1 0-2Zm-7.257 1.914L4 7.172l1.414-1.415l2.329 2.329L12.828 3l1.415 1.414l-6.5 6.5Z"/></g></g>`),
		g.Group(children),
	)
}
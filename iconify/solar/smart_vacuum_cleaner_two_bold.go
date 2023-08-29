package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartVacuumCleanerTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 6.5a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5Z"/><path fill-rule="evenodd" d="M18.52 19.332A9.96 9.96 0 0 1 12 21.75a9.96 9.96 0 0 1-6.52-2.418L2.53 22.28a.75.75 0 0 1-1.06-1.06l2.948-2.95A9.96 9.96 0 0 1 2 11.75c0-5.523 4.477-10 10-10s10 4.477 10 10a9.96 9.96 0 0 1-2.418 6.52l2.948 2.95a.75.75 0 1 1-1.06 1.06l-2.95-2.948ZM8.25 8.75a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Zm4.5 7a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
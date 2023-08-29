package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartVacuumCleanerTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m6.06 19l-3.53 3.53a.75.75 0 0 1-1.06-1.06L5 17.94L6.06 19Zm11.88 0l3.53 3.53a.75.75 0 1 0 1.06-1.06L19 17.94L17.94 19Z" clip-rule="evenodd" opacity=".5"/><path d="M9.75 8.75a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z"/><path fill-rule="evenodd" d="M12 21.5c5.523 0 10-4.477 10-10s-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10ZM8.25 8.75a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Zm4.5 7a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
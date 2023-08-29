package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 3.5a2 2 0 0 1 2-2h10.586a2 2 0 0 1 1.414.586L17.914 4.5a2 2 0 0 1 .586 1.414V16.5a2 2 0 0 1-2 2h-13a2 2 0 0 1-2-2v-13Zm2 0v13h13V5.914L14.086 3.5H3.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 11a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7h-2v-7H7v7H5v-7Zm.5-6a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V3a1 1 0 1 0-2 0v2h-4V3a1 1 0 0 0-2 0v2Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}
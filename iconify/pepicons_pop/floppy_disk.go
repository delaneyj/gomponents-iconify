package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1.5 3.5a2 2 0 0 1 2-2h10.586a2 2 0 0 1 1.414.586L17.914 4.5a2 2 0 0 1 .586 1.414V16.5a2 2 0 0 1-2 2h-13a2 2 0 0 1-2-2v-13Zm2 0v13h13V5.914L14.086 3.5H3.5Z"/><path d="M5 11a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7h-2v-7H7v7H5v-7Zm.5-6a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V3a1 1 0 1 0-2 0v2h-4V3a1 1 0 0 0-2 0v2Z"/></g>`),
		g.Group(children),
	)
}
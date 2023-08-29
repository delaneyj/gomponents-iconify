package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFloppyDiskCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.5 6.5a2 2 0 0 1 2-2h10.586a2 2 0 0 1 1.414.586L20.914 7.5a2 2 0 0 1 .586 1.414V19.5a2 2 0 0 1-2 2h-13a2 2 0 0 1-2-2v-13Zm2 0v13h13V8.914L17.086 6.5H6.5Z"/><path d="M8 14a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7h-2v-7h-6v7H8v-7Zm.5-6a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V6a1 1 0 1 0-2 0v2h-4V6a1 1 0 0 0-2 0v2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFloppyDiskCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
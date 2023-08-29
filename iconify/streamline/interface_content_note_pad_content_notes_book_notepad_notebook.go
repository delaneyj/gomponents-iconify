package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentNotePadContentNotesBookNotepadNotebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 3.5v-3m3 3v-3m3 3v-3"/><rect width="13" height="11.5" x=".5" y="2" rx="1"/></g>`),
		g.Group(children),
	)
}
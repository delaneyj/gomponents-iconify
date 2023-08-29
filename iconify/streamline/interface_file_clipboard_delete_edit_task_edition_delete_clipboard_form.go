package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFileClipboardDeleteEditTaskEditionDeleteClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/><path d="m5.5 6.5l3 3m0-3l-3 3"/></g>`),
		g.Group(children),
	)
}
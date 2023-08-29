package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFileClipboardBlockEditTaskEditionBlockClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 13.5h-4a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1H3m5 0h1.5a1 1 0 0 1 1 1v2"/><rect width="5" height="2.5" x="3" y=".5" rx="1"/><circle cx="10.25" cy="10.25" r="3.25"/><path d="m7.95 12.55l4.6-4.6"/></g>`),
		g.Group(children),
	)
}
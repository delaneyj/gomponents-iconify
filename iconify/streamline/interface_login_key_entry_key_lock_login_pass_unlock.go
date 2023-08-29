package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLoginKeyEntryKeyLockLoginPassUnlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.62 7.38L11.5 1.5l2 2m-4.25.25L11 5.5"/><circle cx="3.5" cy="9.5" r="3"/></g>`),
		g.Group(children),
	)
}
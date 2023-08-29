package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditSkullOneCrashDeathDeleteDieErrorGarbageRemoveSkullTrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13 6.5a6 6 0 1 0-9.5 4.87v1.13a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-1.13A6 6 0 0 0 13 6.5Z"/><circle cx="4.5" cy="7" r=".5"/><circle cx="9.5" cy="7" r=".5"/><path d="M6 11.5v2m2-2v2"/></g>`),
		g.Group(children),
	)
}
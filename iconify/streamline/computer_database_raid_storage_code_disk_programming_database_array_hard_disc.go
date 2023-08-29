package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDatabaseRaidStorageCodeDiskProgrammingDatabaseArrayHardDisc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="7" cy="3" rx="6.5" ry="2.5"/><path d="M.5 3v8c0 1.38 2.91 2.5 6.5 2.5s6.5-1.12 6.5-2.5V3"/><path d="M13.5 7c0 1.38-2.91 2.5-6.5 2.5S.5 8.38.5 7"/></g>`),
		g.Group(children),
	)
}
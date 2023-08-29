package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerLogoAndroidAndroidCodeAppsBugdroidProgramming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 13.5v-8a4 4 0 0 1 8 0v8M3 11h8"/><path d="M.5 10V8a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v2M11 .5L9.28 2.22M3 .5l1.72 1.72"/></g>`),
		g.Group(children),
	)
}
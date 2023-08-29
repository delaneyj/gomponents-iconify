package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingScriptFileCodeTwoCodeProgrammingTerminalShellFileLineCommandFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 13.5l2-2l-2-2m2-2v-6a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1H10"/><path d="M8.5.5v5h5m-9 8h3"/></g>`),
		g.Group(children),
	)
}
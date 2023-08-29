package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingScriptFileCodeOneCodeFilesAngleProgrammingFileBracket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 13.5l-2-2l2-2m3 4l2-2l-2-2"/><path d="M2.5 7.5v-6a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1h-3"/><path d="M8.5.5v5h5"/></g>`),
		g.Group(children),
	)
}
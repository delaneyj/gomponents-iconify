package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingScriptHtmlFiveLanguageFiveCodeProgrammingHtml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12 10l-5 3.5L2 10L.5.5h13L12 10z"/><path d="M9.5 3.5h-5L5 6h4v2.5L7 10L5 8.5"/></g>`),
		g.Group(children),
	)
}
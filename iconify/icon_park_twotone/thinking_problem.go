package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinkingProblem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTThinkingProblem0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m38 21l5 9l-5 1v6h-3l-6-1l-1 7H13l-2-10.381C7.92 29.703 6 25.576 6 21c0-8.837 7.163-16 16-16s16 7.163 16 16Z"/><path d="M17 19a5 5 0 1 1 5 5v3m0 6v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTThinkingProblem0)"/>`),
		g.Group(children),
	)
}
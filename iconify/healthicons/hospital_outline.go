package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M20 20h-4v3h4v-3Zm-4 5h4v3h-4v-3Zm4 5h-4v3h4v-3Zm2-10h4v3h-4v-3Zm4 5h-4v3h4v-3Zm-4 5h4v3h-4v-3Zm10-10h-4v3h4v-3Zm-4 5h4v3h-4v-3Zm4 5h-4v3h4v-3Zm-7-15v-3h3v-2h-3V7h-2v3h-3v2h3v3h2Z"/><path fill-rule="evenodd" d="M17 6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2h8v2h-2v34h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2V8h-1V6h7Zm0 5h-4v31h4v-4h-1v-2h16v2h-1v4h4V11h-4v5a2 2 0 0 1-2 2H19a2 2 0 0 1-2-2v-5Zm0-2h-4V8h4v1Zm2-3h10v10H19V6Zm4 36h-4v-4h4v4Zm6 0v-4h-4v4h4Zm6-33V8h-4v1h4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
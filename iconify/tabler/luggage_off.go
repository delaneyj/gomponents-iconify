package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6h6a2 2 0 0 1 2 2v6m0 4a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V8c0-.546.218-1.04.573-1.4M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v1m-9 4h4m4 0h4M6 16h10m-7 4v1m6-1v1M3 3l18 18"/>`),
		g.Group(children),
	)
}
package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiSecureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M15 26h18v-2H15v2Zm18 5H15v-2h18v2Zm-18 5h18v-2H15v2Z"/><path fill-rule="evenodd" d="M39 20a2 2 0 0 0-2-2h-2v-1c0-6.075-4.925-11-11-11s-11 4.925-11 11v1h-2a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h26a2 2 0 0 0 2-2V20Zm-24-3a9 9 0 1 1 18 0v1h-2v-1a7 7 0 1 0-14 0v3h20v20H11V20h4v-3Zm9-5a5 5 0 0 1 5 5v1H19v-1a5 5 0 0 1 5-5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MammogramStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 16h-2V4H16V2h12a2.002 2.002 0 0 1 2 2Z"/><path fill="currentColor" d="M25 23h-2V9H9V7h14a2.002 2.002 0 0 1 2 2Z"/><path fill="currentColor" d="M18 30H4a2.002 2.002 0 0 1-2-2V14a2.002 2.002 0 0 1 2-2h14a2.002 2.002 0 0 1 2 2v14a2.003 2.003 0 0 1-2 2ZM4 14v14h14.002L18 14Z"/><path fill="currentColor" d="M13 18h1v-2h-1a5.008 5.008 0 0 0-4.899 4H7v2h1.101A5.008 5.008 0 0 0 13 26h1v-2h-1a3 3 0 0 1 0-6Z"/>`),
		g.Group(children),
	)
}
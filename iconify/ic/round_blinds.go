package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundBlinds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 19V5c0-1.1-.9-2-2-2H6c-1.1 0-2 .9-2 2v14H3c-.55 0-1 .45-1 1s.45 1 1 1h18c.55 0 1-.45 1-1s-.45-1-1-1h-1zM16 9h2v2h-2V9zm-2 2H6V9h8v2zm4-4h-2V5h2v2zm-4-2v2H6V5h8zM6 19v-6h8v1.82A1.746 1.746 0 0 0 15 18a1.746 1.746 0 0 0 1-3.18V13h2v6H6z"/>`),
		g.Group(children),
	)
}
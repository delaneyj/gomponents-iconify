package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TajMahal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M16 28H8a2 2 0 0 0-2 2v14m26-16h8a2 2 0 0 1 2 2v14M20 24h8s5.165-5.678 4-9c-.755-2.15-2.439-3.34-4-5c-1.563-1.66-4-4-4-4s-2.439 2.34-4 4c-1.563 1.66-3.246 2.85-4 5c-1.166 3.322 4 9 4 9Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40M16 24h16v20H16V24Z"/><path stroke-linecap="round" d="M24 34v10m0-40v3M10 24v4m28-4v4M20 44h8"/></g>`),
		g.Group(children),
	)
}
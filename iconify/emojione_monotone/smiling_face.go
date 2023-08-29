package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5c15.163 0 27.5 12.336 27.5 27.5S47.163 59.5 32 59.5z"/><circle cx="20.5" cy="24.592" r="5" fill="currentColor"/><circle cx="43.5" cy="24.592" r="5" fill="currentColor"/><path fill="currentColor" d="M48.11 37.02c-4.328 6.107-9.451 7.644-16.111 7.644c-6.659 0-11.782-1.536-16.11-7.644c-.603-.85-2.19-.315-1.839.919c2.273 8.005 10 12.667 17.95 12.667c7.95 0 15.678-4.662 17.951-12.667c.349-1.235-1.238-1.769-1.841-.919"/>`),
		g.Group(children),
	)
}
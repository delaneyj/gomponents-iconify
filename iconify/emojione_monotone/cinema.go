package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M35 27.869A7.994 7.994 0 0 1 28 32h14a7.994 7.994 0 0 1-7-4.131"/><circle cx="28" cy="24" r="3" fill="currentColor"/><circle cx="42" cy="24" r="3" fill="currentColor"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m14 30c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H23c-1.1 0-2-.9-2-2v-.5L16 48V32l5 2.5V34c0-1.1.9-2 2-2h5a8 8 0 0 1 0-16a7.994 7.994 0 0 1 7 4.131A7.994 7.994 0 0 1 42 16a8 8 0 0 1 0 16h4"/>`),
		g.Group(children),
	)
}
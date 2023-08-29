package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.569 13.432 30 30 30s30-13.432 30-30C62 15.432 48.568 2 32 2zm0 58C16.537 60 4 47.465 4 32C4 16.536 16.537 4 32 4c15.465 0 28 12.536 28 28c0 15.465-12.535 28-28 28z"/>`),
		g.Group(children),
	)
}
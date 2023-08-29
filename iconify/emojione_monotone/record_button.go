package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.569 13.432 30 30 30s30-13.431 30-30C62 15.432 48.568 2 32 2m0 45c-8.283 0-15-6.716-15-15s6.717-15 15-15c8.285 0 15 6.716 15 15s-6.715 15-15 15"/>`),
		g.Group(children),
	)
}
package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm4 30c0 1.477-.81 2.752-2 3.445V38h-3.229l-6.187 10L21 46.068l7.514-12.146A3.937 3.937 0 0 1 28 32c0-1.477.81-2.752 2-3.445V6h4v20.315l3 1.614l-1.427 2.306A3.96 3.96 0 0 1 36 32z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}
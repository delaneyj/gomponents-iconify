package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m5.975 33.5l-3.029-.813c-.273.3-.592.553-.945.758V58h-4V35.445a3.986 3.986 0 0 1-1.905-2.597L15 29.334l1.027-3.834l13.276 3.563c.212-.195.445-.362.696-.508V26h4v2.554a3.972 3.972 0 0 1 1.798 2.251l3.203.859l-1.025 3.836"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}
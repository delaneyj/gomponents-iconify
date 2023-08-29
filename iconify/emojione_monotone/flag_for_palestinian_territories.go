package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPalestinianTerritories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM20.557 41.833l-10.329 7.746a28.378 28.378 0 0 1-1.2-1.6L30.334 32L9.027 16.021c.382-.548.781-1.081 1.2-1.6l10.328 7.746h37.657C59.365 25.229 60 28.541 60 32s-.635 6.771-1.787 9.833H20.557z"/>`),
		g.Group(children),
	)
}
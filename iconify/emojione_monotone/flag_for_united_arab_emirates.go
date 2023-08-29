package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForUnitedArabEmirates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-9 40v16.564a28.288 28.288 0 0 1-2-.746V6.182c.653-.275 1.322-.52 2-.746V22h35.213C59.365 25.062 60 28.541 60 32c0 3.459-.635 6.938-1.787 10H23z"/>`),
		g.Group(children),
	)
}
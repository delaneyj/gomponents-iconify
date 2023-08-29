package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForArmenia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM5.787 22.167h52.426c.246.655.467 1.323.664 2H5.123c.198-.677.418-1.345.664-2zM32 60c-11.601 0-21.575-7.092-25.818-17.167h51.636C53.574 52.908 43.601 60 32 60z"/>`),
		g.Group(children),
	)
}
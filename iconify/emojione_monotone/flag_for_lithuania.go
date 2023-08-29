package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForLithuania(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM6.254 43a27.34 27.34 0 0 1-.679-1.756h52.85A27.943 27.943 0 0 1 57.746 43H6.254zM32 4c11.917 0 22.112 7.486 26.147 18H5.853C9.888 11.486 20.083 4 32 4z"/>`),
		g.Group(children),
	)
}
package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForJamaica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-.001 2c6.532 0 12.724 2.265 17.668 6.422L32 24.014L14.332 10.421C19.275 6.265 25.467 4 31.999 4zm0 56c-6.532 0-12.724-2.265-17.667-6.421L32 39.986l17.666 13.592C44.723 57.735 38.531 60 31.999 60zm19.56-7.99L32 36.963L12.44 52.01a28.176 28.176 0 0 1-4.802-6.232L25.547 32L7.638 18.224a28.11 28.11 0 0 1 4.803-6.233L32 27.037l19.559-15.048a28.193 28.193 0 0 1 4.803 6.233L38.451 32l17.91 13.777a28.135 28.135 0 0 1-4.802 6.233z"/>`),
		g.Group(children),
	)
}
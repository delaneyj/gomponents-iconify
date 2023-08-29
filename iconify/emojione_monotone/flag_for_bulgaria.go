package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBulgaria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m0 2c11.917 0 22.112 7.486 26.147 18H5.853C9.888 11.486 20.083 4 32 4M6.254 43a27.929 27.929 0 0 1-.762-2h53.016a27.46 27.46 0 0 1-.762 2H6.254"/>`),
		g.Group(children),
	)
}
package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M74.055 143.173L0 325.823h53.607l13.263-31.777h76.266l14.645 31.778h54.436l-82.345-182.65H74.055zm27.909 51.673l25.974 62.726H78.752l23.212-62.726zm55.748-51.673L258.5 368.827l104.105-225.654h-55.057L258.5 253.703l-46.422-110.53h-54.367zm302.367 0h-71.568l-80.962 182.65h54.988l29.221-65.557h12.573l17.34 65.558h56.37l-25.422-72.12c76.015-10.86 79.64-107.371 7.46-110.53zm-30.672 78.753h-24.005l19.584-43.66h19.62c25.173 3.84 20.152 42.054-15.199 43.66z"/>`),
		g.Group(children),
	)
}
package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M230 209.2L32 405.58L106.65 480l197.59-198.17c46.47 17.46 105.52 12.54 143-24.78c40.44-40.32 40.35-108 16.81-156.79l-87.33 87.06l-52.32-52.13l87.33-87.06C363 24.46 294.67 24.34 254.23 64.66c-38.03 37.91-42.78 97.6-24.23 144.54Z"/>`),
		g.Group(children),
	)
}
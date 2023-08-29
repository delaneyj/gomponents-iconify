package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inspire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M32 32v448h448V32Zm77.673 79.678a34.076 34.076 0 0 1 34.076 34.076a34.076 34.076 0 0 1-34.076 34.076a34.076 34.076 0 0 1-34.076-34.076a34.076 34.076 0 0 1 34.076-34.076zm76.17 9.02l64.173.16l116.23 187.258V120.698h59.132v280.626h-65.145l-112.25-185.413v185.413h-62.14ZM78.604 203.884h62.139v197.44H78.604z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentificationCardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 108a12 12 0 0 1 12-12h32a12 12 0 0 1 0 24h-32a12 12 0 0 1-12-12Zm12 52h32a12 12 0 0 0 0-24h-32a12 12 0 0 0 0 24Zm80-104v144a20 20 0 0 1-20 20H40a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h176a20 20 0 0 1 20 20Zm-24 4H44v136h168ZM54.28 159.37A43.82 43.82 0 0 1 67.53 142a36 36 0 1 1 56.94 0a43.84 43.84 0 0 1 13.26 17.37a12 12 0 0 1-22.15 9.26C112.48 161.19 104.42 156 96 156s-16.47 5.2-19.59 12.63a12 12 0 1 1-22.13-9.26ZM84 120a12 12 0 1 0 12-12a12 12 0 0 0-12 12Z"/>`),
		g.Group(children),
	)
}
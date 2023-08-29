package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftPowerpointLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M96 100H80a4 4 0 0 0-4 4v48a4 4 0 0 0 8 0v-12h12a20 20 0 0 0 0-40Zm0 32H84v-24h12a12 12 0 0 1 0 24Zm40-104a100.3 100.3 0 0 0-80 40H40a12 12 0 0 0-12 12v96a12 12 0 0 0 12 12h16a100 100 0 1 0 80-160Zm91.91 96H156V80a12 12 0 0 0-12-12h-4V36.09A92.13 92.13 0 0 1 227.91 124ZM132 36.1V68H66.26A92.36 92.36 0 0 1 132 36.1ZM36 176V80a4 4 0 0 1 4-4h104a4 4 0 0 1 4 4v96a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4Zm30.26 12H132v31.9A92.36 92.36 0 0 1 66.26 188ZM140 219.91V188h4a12 12 0 0 0 12-12v-44h71.91A92.13 92.13 0 0 1 140 219.91Z"/>`),
		g.Group(children),
	)
}
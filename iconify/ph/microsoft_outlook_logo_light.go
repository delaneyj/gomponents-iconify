package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftOutlookLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M118 128a30 30 0 1 0-30 30a30 30 0 0 0 30-30Zm-48 0a18 18 0 1 1 18 18a18 18 0 0 1-18-18Zm146-22h-10V48a14 14 0 0 0-14-14h-80a14 14 0 0 0-14 14v18H40a14 14 0 0 0-14 14v96a14 14 0 0 0 14 14h34v18a14 14 0 0 0 14 14h128a14 14 0 0 0 14-14v-88a14 14 0 0 0-14-14Zm-53.75 58L218 123.73v80.54ZM110 48a2 2 0 0 1 2-2h80a2 2 0 0 1 2 2v78.27l-42 30.33l-2-1.45V80a14 14 0 0 0-14-14h-26ZM38 176V80a2 2 0 0 1 2-2h96a2 2 0 0 1 2 2v96a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2Zm48 32v-18h50a14 14 0 0 0 14-14v-6l55.44 40H88a2 2 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}
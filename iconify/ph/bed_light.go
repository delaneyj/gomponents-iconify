package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 74H22V48a6 6 0 0 0-12 0v160a6 6 0 0 0 12 0v-34h212v34a6 6 0 0 0 12 0v-96a38 38 0 0 0-38-38ZM22 86h76v76H22Zm88 76V86h98a26 26 0 0 1 26 26v50Z"/>`),
		g.Group(children),
	)
}
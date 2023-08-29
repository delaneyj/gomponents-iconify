package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247.5 0L34.2 213.3v128l170.6-170.6V512h85.4V170.7l170.6 170.6v-128z"/>`),
		g.Group(children),
	)
}
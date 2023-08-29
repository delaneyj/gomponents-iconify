package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUpTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 24L24 6l19 18H31v18H17V24H5Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUpTwo0)"/>`),
		g.Group(children),
	)
}
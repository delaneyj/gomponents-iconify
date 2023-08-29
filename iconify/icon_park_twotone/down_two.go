package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDownTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m5 24l19 18l19-18H31V6H17v18H5Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDownTwo0)"/>`),
		g.Group(children),
	)
}
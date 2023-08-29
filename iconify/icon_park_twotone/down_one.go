package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDownOne0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M36 19L24 31L12 19h24Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDownOne0)"/>`),
		g.Group(children),
	)
}
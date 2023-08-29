package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Renault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRenault0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M30 4H18L8 24h9l5-10h4l5 10h9L30 4Zm0 40H18L8 24h9l5 10h4l5-10h9L30 44Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRenault0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTheme0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M18 6a6 6 0 0 0 12 0h5.455L42 15.818l-5.727 4.91V42H11.727V20.727L6 15.818L12.546 6H18Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTheme0)"/>`),
		g.Group(children),
	)
}
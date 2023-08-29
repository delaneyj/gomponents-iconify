package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlockNine0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 14h12v12H18V14Zm0 12h12v12H18V26Zm12-12h12v12H30V14ZM6 26h12v12H6V26Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlockNine0)"/>`),
		g.Group(children),
	)
}
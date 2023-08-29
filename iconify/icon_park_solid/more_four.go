package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMoreFour0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 11l7.071-7.071L31.142 11l-7.07 7.071zm13 13l7.071-7.071L44.142 24l-7.07 7.071zM4 24l7.071-7.071L18.142 24l-7.07 7.071zm13 13l7.071-7.071L31.142 37l-7.07 7.071z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMoreFour0)"/>`),
		g.Group(children),
	)
}
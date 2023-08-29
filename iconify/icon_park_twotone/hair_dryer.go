package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HairDryer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHairDryer0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m19.193 21.544l2.319 18.552a3.473 3.473 0 0 1-6.892.862l-2.374-18.989"/><path fill="#555" d="M13 4a9 9 0 0 0 0 18c1.578 0 3.74-.175 6.193-.456l12.403-2.022L44 17.5v-9L28.5 6.25L13 4Z"/><path d="M37 8.2v9.6m7-.3l-12.403 2.022M44 8.5L28.5 6.25"/><path fill="#555" d="M16 13a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHairDryer0)"/>`),
		g.Group(children),
	)
}
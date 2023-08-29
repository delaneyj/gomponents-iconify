package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjacentItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAdjacentItem0"><g fill="none"><path fill="#fff" d="M14 29h28v12H14V29Zm0-22h28v12H14V7Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 13v6h28V7H14v6Zm0 0H6v22h8m0 0v6h28V29H14v6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 13H6v22h8"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 29h28v12H14V29Zm0-22h28v12H14V7Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAdjacentItem0)"/>`),
		g.Group(children),
	)
}
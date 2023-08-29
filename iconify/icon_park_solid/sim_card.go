package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSimCard0"><g fill="none"><path stroke="#fff" stroke-width="4" d="M10 44h28a2 2 0 0 0 2-2V14.885a2 2 0 0 0-.655-1.48L29.572 4.52A2 2 0 0 0 28.227 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><circle cx="17" cy="14" r="3" fill="#fff"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 23h16v14H16z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSimCard0)"/>`),
		g.Group(children),
	)
}
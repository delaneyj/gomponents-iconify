package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 11.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.5 9.286c0 2.673 2.653 7.214 4.5 7.214c1.848 0 4.5-4.541 4.5-7.214C17.5 6.65 15.493 4.5 13 4.5S8.5 6.65 8.5 9.286Zm8 0c0 2.193-2.348 6.214-3.5 6.214c-1.151 0-3.5-4.02-3.5-6.214C9.5 7.187 11.075 5.5 13 5.5s3.5 1.687 3.5 3.786Z" clip-rule="evenodd"/><path d="M16.435 12.14a.5.5 0 0 1 .369-.929a3 3 0 0 1 1.74 1.84l1.334 4A3 3 0 0 1 17.03 21H8.97a3 3 0 0 1-2.846-3.949l1.333-4a3 3 0 0 1 1.783-1.857a.5.5 0 1 1 .355.935a2 2 0 0 0-1.19 1.239l-1.333 4A2 2 0 0 0 8.97 20h8.062a2 2 0 0 0 1.897-2.633l-1.332-4a2 2 0 0 0-1.16-1.226Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}